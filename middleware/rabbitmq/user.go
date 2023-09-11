package rabbitmq

import (
	"context"
	"douyin-server/database/dao"
	"douyin-server/middleware/gorse"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type UserDealer struct {
}

type RegisterMessage struct {
	Filename string `json:"id"`
	Token    string `json:"token"`
	Title    string `json:"title"`
}

func (dealer UserDealer) DealWith(msg <-chan amqp.Delivery) (err error) {
	for d := range msg {
		// 参数解析
		user := &dao.Duser{}
		if err = json.Unmarshal(d.Body, user); err != nil {
			log.Println("json解析失败")
		}
		if err = gorse.RegisterToGorse(context.Background(), user); err != nil {
			log.Println("RegisterToGorse Err : ", err)
		}
	}
	return
}
