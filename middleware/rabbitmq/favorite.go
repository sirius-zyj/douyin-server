package rabbitmq

import (
	"context"
	"douyin-server/database/dao"
	"douyin-server/middleware/gorse"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type FavoriteDealer struct {
}

func (dealer FavoriteDealer) DealWith(msg <-chan amqp.Delivery) (err error) {
	for d := range msg {
		// 参数解析
		fa := &dao.Dfavorite{}
		if err := json.Unmarshal(d.Body, fa); err != nil {
			log.Println("json解析失败")
		}

		if err = gorse.FavoriteToGorse(context.Background(), fa); err != nil {
			log.Println("FavoriteToGorse Err : ", err)
		}
	}
	return nil
}
