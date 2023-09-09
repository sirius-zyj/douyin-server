package rabbitmq

import (
	"douyin-server/database/dao"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type FavoriteDealer struct {
}

func (dealer FavoriteDealer) DealWith(msg <-chan amqp.Delivery) error {
	for d := range msg {
		// 参数解析
		fa := &dao.Dfavorite{}
		if err := json.Unmarshal(d.Body, fa); err != nil {
			log.Println("json解析失败")
		}

		if fa.Action_type == "1" {
			log.Println("点赞成功")
		} else {
			log.Println("取消点赞成功")
		}
	}
	return nil
}
