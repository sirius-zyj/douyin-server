package rabbitmq

import (
	"douyin-server/config"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

const MQURL = "amqp://xwh:123456@58.87.89.6:5672/"

type ServiceDealer interface {
	DealWith(<-chan amqp.Delivery) error
}

type RabbitMQ struct {
	conn  *amqp.Connection
	mqurl string

	channel   *amqp.Channel
	queueName string
	exchange  string
	dealer    ServiceDealer
}

var mq *RabbitMQ

func initRabbitMQ(queueName string) {
	mq = &RabbitMQ{
		mqurl:     MQURL,
		queueName: queueName,
	}
	var err error
	if mq.conn, err = amqp.Dial(mq.mqurl); err != nil {
		log.Panicln("Rabbitmq连接失败")
	}

	if cha, err := mq.conn.Channel(); err != nil {
		log.Panicln(err)
	} else {
		mq.channel = cha
	}
	if _, err := mq.channel.QueueDeclare(
		mq.queueName,
		//是否持久化
		false,
		//是否为自动删除
		true,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	); err != nil {
		log.Panicln(err)
	}

	switch queueName {
	case config.FavoriteServiceName:
		mq.dealer = FavoriteDealer{}
	case config.PublishServiceName:
		mq.dealer = PublishDealer{}
	}

}

// Publish like操作的发布配置。
func (l *RabbitMQ) produce(message interface{}) (err error) {
	jsonMessage, _ := json.Marshal(message)
	if err = l.channel.Publish(
		l.exchange,
		l.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonMessage,
		}); err != nil {
		log.Println(err)
	}
	return
}

// Consumer like关系的消费逻辑。
func (l *RabbitMQ) consumer() {
	//2、接收消息
	messages, err1 := l.channel.Consume(
		l.queueName,
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		//消息队列是否阻塞
		false,
		nil,
	)
	if err1 != nil {
		panic(err1)
	}

	forever := make(chan bool)
	go l.dealer.DealWith(messages)

	log.Printf("[*] Waiting for messagees,To exit press CTRL+C")

	<-forever

}

func CreateRabbitMQ(queueName string) {
	initRabbitMQ(queueName)
	go mq.consumer()
}

func DestoryRabbitMQ() {
	if _, err := mq.channel.QueueDelete(mq.queueName, false, false, false); err != nil {
		log.Println("mq queue删除失败")
	}
	if err := mq.channel.Close(); err != nil {
		log.Println("mq channel 删除失败")
	}
	if err := mq.conn.Close(); err != nil {
		log.Println("mq conn删除失败")
	}
}

func AddToMQ(message interface{}) error {
	return mq.produce(message)
}
