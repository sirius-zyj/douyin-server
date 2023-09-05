package mysql2redis

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"sync"

	"github.com/duanhf2012/origin/log"

	"github.com/siddontang/go/ioutil2"

	"github.com/go-mysql-org/go-mysql/mysql"
)

type Output struct {
	Config *CanalConfig
	//	PosFile         *os.File
	Adapters        map[string]WriteAdapter
	DataChannel     chan *RawLogEntity
	PosChannel      chan *mysql.Position
	Datas           []*RawLogEntity
	lastWriteTime   time.Time
	writeLock       *sync.Mutex
	writeDataLength int64
}

func CreateByName(name string) (*Output, error) {

	currConfig, isConfigExist := GetConfigure().CanalConfigs[name]
	if !isConfigExist {
		return nil, fmt.Errorf("output config is not exist for name %s", name)
	}
	currOutput := &Output{}
	currOutput.Config = currConfig
	currOutput.Adapters = map[string]WriteAdapter{}
	//	currOutput.DataChannel = make(chan *common.RawLogEntity, currConfig.CacheSize)
	currOutput.DataChannel = make(chan *RawLogEntity, currConfig.CacheSize)
	currOutput.PosChannel = make(chan *mysql.Position, currConfig.CacheSize)
	currOutput.Datas = []*RawLogEntity{}
	currOutput.lastWriteTime = time.Now()
	currOutput.writeLock = &sync.Mutex{}
	posPath := path.Dir(currOutput.Config.Posconfigfile)
	makePosPathErr := os.MkdirAll(posPath, os.ModePerm)

	if makePosPathErr != nil {
		log.Error(makePosPathErr.Error())
		panic(makePosPathErr)
	}

	if currConfig.Redis != nil {
		currOutput.Adapters[currConfig.Redis.GetConfigName()] = createAdapter(currConfig.Redis)
	}

	return currOutput, nil
}
func CreateAdapterWithName(conf CommonConfig) (WriteAdapter, error) {
	if conf.GetConfigName() == "Redis" {
		return CreateAdapter(conf.(*RedisConfig)), nil
	}

	log.Error("Config Type %v is not support !!!!", conf.GetConfigName())
	return nil, errors.New("config type error")
}

func createAdapter(conf CommonConfig) WriteAdapter {
	currAdapter, createAdapterErr := CreateAdapterWithName(conf)
	if createAdapterErr != nil {
		log.Error(createAdapterErr.Error())
		panic(createAdapterErr)
	}
	return currAdapter
}

func (output *Output) Run() {
	output.lastWriteTime = time.Now()
	go output.writeTimeProcess()
	// go output.processPos()
	for {
		func() {
			defer func() {
				err := recover()
				if err != nil {
					log.Error("crash, err:%s", err)
				}
			}()

			currData := <-output.DataChannel

			output.Datas = append(output.Datas, currData)

			dataLength := len(output.Datas)
			// log.Debug("Data Length = %v", dataLength)
			if dataLength >= output.Config.Bulksize {
				output.writeDataToAdapter()
			}
		}()
	}
}

func (output *Output) writeDataToAdapter() {
	output.writeLock.Lock()
	defer output.writeLock.Unlock()
	dataLength := len(output.Datas)
	if dataLength > 0 {
		for _, currAdapter := range output.Adapters {
			adapterWriteErr := currAdapter.Write(output.Datas)
			if adapterWriteErr != nil {
				log.Error(adapterWriteErr.Error())
				// panic(adapterWriteErr)
				configTimeDuration := output.Config.Flushbulktime * int64(time.Millisecond)
				time.Sleep(time.Duration(configTimeDuration))
				return
			}
		}

		output.Datas = []*RawLogEntity{}
		output.lastWriteTime = time.Now()
		output.writeDataLength = output.writeDataLength + int64(dataLength)
	}
}

func (output *Output) writeTimeProcess() {
	for {
		func() {
			defer func() {
				err := recover()
				if err != nil {
					log.Error("writeTimeProcess crash, err:%s", err)
				}
			}()

			currTimeDuration := time.Now().UnixNano() - output.lastWriteTime.UnixNano()
			configTimeDuration := output.Config.Flushbulktime * int64(time.Millisecond)
			if currTimeDuration >= configTimeDuration {
				output.writeDataToAdapter()
				time.Sleep(time.Duration(configTimeDuration))
			} else {
				time.Sleep(time.Duration(configTimeDuration - currTimeDuration))
			}
		}()
	}
}

func (output *Output) Write(data *RawLogEntity) {
	output.DataChannel <- data
}

func (output *Output) WritePos(pos *mysql.Position) {
	output.PosChannel <- pos
}

func (output *Output) Empty() bool {
	return len(output.PosChannel) == 0 && len(output.DataChannel) == 0
}

func (output *Output) processPos() {
	for {
		func() {
			defer func() {
				err := recover()
				if err != nil {
					log.Error("processPos crash, err:%s", err)
				}
			}()

			var posData *mysql.Position
			ccap := cap(output.PosChannel)
			for k := 0; k < ccap; k++ {
				select {
				case pos := <-output.PosChannel:
					posData = pos
				default:
					break
				}
			}

			if posData != nil {
				binFileName := fmt.Sprintf("bin_name = \"%v\" \r\n", posData.Name)
				binFilePos := fmt.Sprintf("bin_pos = %v \r\n", posData.Pos)
				content := binFileName + binFilePos
				if output.Config.Posconfigfile != "" {
					if err := ioutil2.WriteFileAtomic(output.Config.Posconfigfile, []byte(content), 0777); err != nil {
						log.Error("canal save master info to file %s err %v", output.Config.Posconfigfile, err)
					}
				} else {
					log.Debug("bin_name = \"%v\" bin_pos = %v", posData.Name, posData.Pos)
				}
			}
			time.Sleep(time.Millisecond * 60)
		}()
	}
}
