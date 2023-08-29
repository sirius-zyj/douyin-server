package mysql2redis

import (
	"fmt"

	"github.com/duanhf2012/origin/log"
)

var UserPrefix = "user{id:%d}" //redis中user的前缀

type UserMicroservice struct {
}

func (microservice *UserMicroservice) DealData(currEntity *RawLogEntity, commands *[]*RedisPipelineCommand) (err error) {

	if currEntity.Action == "update" || currEntity.Action == "delete" {
		flag := false
		if currEntity.Action == "update" {
			flag = true
		}
		if flag && len(currEntity.Rows)%2 != 0 {
			log.Error("update event, row's count must be even")
			return
		}
		for idx, row := range currEntity.Rows {
			if flag && idx%2 == 0 {
				continue
			}
			key := fmt.Sprintf(UserPrefix, row[currEntity.HeaderMap["id"]])
			var currCommand = PipelineCommandPoll.Get().(*RedisPipelineCommand)
			currCommand.CommandName = "DEL"
			currCommand.Key = key
			currCommand.TabName = currEntity.TableName
			currCommand.Args = []interface{}{}
			*commands = append(*commands, currCommand)
		}
	}
	return
}
