package mysql2redis

import (
	"github.com/duanhf2012/origin/log"
	"github.com/duanhf2012/origin/util/sync"
	"github.com/garyburd/redigo/redis"
	// "server/common/settings"
)

type RedisPipelineCommand struct {
	CommandName string
	Key         string
	TabName     string
	Args        []interface{}

	ref bool
}

func (slf *RedisPipelineCommand) Reset() {

}

func (slf *RedisPipelineCommand) IsRef() bool {
	return slf.ref
}

func (slf *RedisPipelineCommand) Ref() {
	slf.ref = true
}

func (slf *RedisPipelineCommand) UnRef() {
	slf.ref = false
}

var PipelineCommandPoll = sync.NewPoolEx(make(chan sync.IPoolData, 1000), func() sync.IPoolData {
	return &RedisPipelineCommand{}
})

type RedisAdapter struct {
	WriteAdapter
	redisClient map[string]redis.Conn
	Config      *RedisConfig
}

func CreateAdapter(conf *RedisConfig) WriteAdapter {
	adapter := &RedisAdapter{Config: conf}
	log.Debug("conf's table count:%d", len(conf.Tables))
	adapter.redisClient = make(map[string]redis.Conn, len(conf.Tables))
	for n, cfg := range conf.Tables {
		//adapter.redisClient[n] = {}
		if conf.Password != "" {
			option := redis.DialPassword(conf.Password)
			currClient, redisConnErr := redis.Dial("tcp", conf.Address, option)
			if redisConnErr != nil {
				//log4go.Error(redisConnErr)
				panic(redisConnErr)
			}

			currClient.Do("SELECT", cfg.DB)
			adapter.redisClient[n] = currClient
		} else {
			currClient, redisConnErr := redis.Dial("tcp", conf.Address)
			if redisConnErr != nil {
				//log4go.Error(redisConnErr)
				panic(redisConnErr)
			}

			currClient.Do("SELECT", cfg.DB)
			adapter.redisClient[n] = currClient
		}
	}

	return adapter
}

func (redisAdapter *RedisAdapter) Write(entities []*RawLogEntity) error {
	commands := make([]*RedisPipelineCommand, 0)
	for _, currEntity := range entities {

		tabName := currEntity.TableName
		if _, exist := redisAdapter.Config.Tables[tabName]; !exist {
			continue
		}

		if microservice := createMicroservice(tabName); microservice != nil {
			microservice.DealData(currEntity, &commands)
			RawLogEntityPoll.Put(currEntity)
		}
	}

	_, commandsSendErrors := redisAdapter.SendPipelineCommands(commands)
	for _, currErr := range commandsSendErrors {
		if currErr != nil {
			return currErr
		}
	}

	return nil
}

func (redisAdapter *RedisAdapter) Close() error {
	for _, conn := range redisAdapter.redisClient {
		closeErr := conn.Close()
		if closeErr != nil {
			log.Error(closeErr.Error())
			return closeErr

		}
	}
	return nil
}

func (redisAdapter *RedisAdapter) SendPipelineCommands(commands []*RedisPipelineCommand) ([]interface{}, []error) {
	errorList := make([]error, 0, len(commands)+1)

	cotMap := make(map[string]int)
	for name, _ := range redisAdapter.redisClient {
		cotMap[name] = 0
	}

	for _, cmd := range commands {
		client := redisAdapter.redisClient[cmd.TabName]

		currErr := client.Send(cmd.CommandName, append([](interface{}){cmd.Key}, cmd.Args...)...)
		// client.Send("EXPIRE", cmd.Key, settings.Settings.Global.RedisExpire)
		// client.Send("EXPIRE", cmd.Key, 120)

		cotMap[cmd.TabName] += 1
		if currErr != nil {
			log.Error(currErr.Error())
			errorList = append(errorList, currErr)
		}
		PipelineCommandPoll.Put(cmd)
	}

	// log.Debug("Send finished!!")

	for _, client := range redisAdapter.redisClient {
		fulshErr := client.Flush()
		if fulshErr != nil {
			log.Error(fulshErr.Error())
			errorList = append(errorList, fulshErr)

			return nil, errorList
		}
	}

	// log.Debug("Flush finished!!")

	replys := [](interface{}){}

	for name, cot := range cotMap {
		replysLength := cot

		for i := 0; i < replysLength; i++ {
			reply, receiveErr := redisAdapter.redisClient[name].Receive()

			if receiveErr != nil {
				log.Error(receiveErr.Error())
				errorList = append(errorList, receiveErr)
			}

			replys = append(replys, reply)
		}
	}

	// log.Debug("Receive finished!!")

	if len(errorList) != 0 {
		return replys, errorList
	}

	return replys, nil
}

type Microservice interface {
	DealData(*RawLogEntity, *[]*RedisPipelineCommand) error
}

func createMicroservice(name string) Microservice {
	if name == "videos" {
		return &VideoMicroservice{}
	} else if name == "users" {
		return &UserMicroservice{}
	} else if name == "favorite" {
		return &FavoriteMicroservice{}
	} else if name == "comments" {
		return &CommentMicroservice{}
	} else if name == "follows" {
		return &FollowMicroservice{}
	}
	return nil
}
