package main

import (
	"douyin-server/dao"
	feed "douyin-server/rpc/kitex_gen/feed/feedservice"
	"log"
)

func main() {
	svr := feed.NewServer(new(FeedServiceImpl))

	dao.Init()

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
