package mysql2redis

import (
	"fmt"

	"github.com/duanhf2012/origin/log"
)

var VideoPrefix = "feed{id:%d}"             //redis中video的前缀
var PublishPrefix = "publish{author_id:%d}" //redis中publish的前缀

type VideoMicroservice struct {
}

func (microservice *VideoMicroservice) DealData(currEntity *RawLogEntity, commands *[]*RedisPipelineCommand) (err error) {

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
			key := fmt.Sprintf(VideoPrefix, row[currEntity.HeaderMap["id"]])
			var currCommand = PipelineCommandPoll.Get().(*RedisPipelineCommand)
			currCommand.CommandName = "DEL"
			currCommand.Key = key
			currCommand.TabName = currEntity.TableName
			currCommand.Args = []interface{}{}
			*commands = append(*commands, currCommand)
		}
	} else if currEntity.Action == "insert" {
		for _, row := range currEntity.Rows {
			key := fmt.Sprintf(PublishPrefix, row[currEntity.HeaderMap["author_id"]])
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
