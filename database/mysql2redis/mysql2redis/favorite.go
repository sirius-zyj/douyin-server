package mysql2redis

import (
	"fmt"

	"github.com/duanhf2012/origin/log"
)

var FavoriteDataPrefix = "favorite{user_id:%d; video_id:%d}" //redis中follow的前缀

type FavoriteMicroservice struct {
}

func (microservice *FavoriteMicroservice) DealData(currEntity *RawLogEntity, commands *[]*RedisPipelineCommand) (err error) {

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
			key := fmt.Sprintf(FavoriteDataPrefix, row[currEntity.HeaderMap["user_id"]], row[currEntity.HeaderMap["video_id"]])
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
