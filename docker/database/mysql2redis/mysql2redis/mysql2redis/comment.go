package mysql2redis

import (
	"fmt"

	"github.com/duanhf2012/origin/log"
)

var CommentPrefix = "comment{video_id:%d}" //redis中评论的前缀

type CommentMicroservice struct {
}

func (microservice *CommentMicroservice) DealData(currEntity *RawLogEntity, commands *[]*RedisPipelineCommand) (err error) {
	if currEntity.Action == "update" || currEntity.Action == "delete" || currEntity.Action == "insert" {
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
			key := fmt.Sprintf(CommentPrefix, row[currEntity.HeaderMap["video_id"]])
			var currCommand = PipelineCommandPoll.Get().(*RedisPipelineCommand)
			currCommand.CommandName = "DEL"
			currCommand.Key = key
			currCommand.TabName = currEntity.TableName
			currCommand.Args = []interface{}{}
			log.Debug("CommentMicroservice DealData command:%v", *currCommand)
			*commands = append(*commands, currCommand)
		}
	}
	return
}
