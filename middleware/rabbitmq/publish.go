package rabbitmq

import (
	"context"
	"douyin-server/database/dao"
	"douyin-server/middleware/gorse"
	"douyin-server/middleware/jwt"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

type PublishDealer struct {
}

type PublishActionMessage struct {
	Filename string `json:"id"`
	Token    string `json:"token"`
	Title    string `json:"title"`
}

func (dealer PublishDealer) DealWith(msg <-chan amqp.Delivery) error {
	for d := range msg {
		// 参数解析
		req := &PublishActionMessage{}
		if err := json.Unmarshal(d.Body, req); err != nil {
			log.Println("json解析失败")
		}
		defer os.Remove(req.Filename)
		data, _ := ioutil.ReadFile(req.Filename)

		token, video_title := req.Token, req.Title
		if playUrl, coverUrl, err := dao.UploadVideo(&data); err == nil {
			//------创建视频--------
			userId := jwt.GetUserIdByToken(token)

			video := dao.Dvideo{
				Id:          dao.SnowFlakeNode.Generate().Int64(),
				Author_id:   userId,
				Play_url:    playUrl,
				Cover_url:   coverUrl,
				Upload_time: time.Now(),
				Title:       video_title,
			}
			if err = dao.Tran_InsertVideo(video); err != nil {
				log.Println("Insert_error: ", err)
				return err
			}

			log.Println("上传视频成功, title: ", req.Title)

			//------创建视频推荐--------
			if err = gorse.PublishToGorse(context.Background(), &video); err != nil {
				log.Println("推荐视频失败")
			}
		} else {
			log.Println("上传视频失败")
		}
	}
	return nil
}
