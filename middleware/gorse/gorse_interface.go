package gorse

import (
	"context"
	"douyin-server/config"
	"douyin-server/database/dao"
	"douyin-server/middleware/wxyy"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	gorseInstance *GorseClient
	once          sync.Once
)

// GetInstance 返回单例对象的实例
func GetInstance() *GorseClient {
	once.Do(func() {
		gorseInstance = NewGorseClient(config.GorseAddr, "")
		wxyy.GetInstance()
	})
	return gorseInstance
}

func PublishToGorse(ctx context.Context, video *dao.Dvideo) (err error) {
	item := &Item{
		ItemId:    strconv.FormatInt(video.Id, 10),
		IsHidden:  false,
		Timestamp: video.Upload_time.Format("2006-01-02"),
	}
	defer func(it *Item) {
		if _, err = GetInstance().InsertItem(ctx, *it); err != nil {
			log.Println("Insert Err", err)
			return
		}
	}(item)

	text, err := wxyy.GetInstance().Video2Txt(video.Play_url)
	if err != nil || text == nil {
		return
	}
	keyword, err := wxyy.GetInstance().Txt2Keyword(text)
	if err != nil {
		return
	}
	for _, v := range keyword {
		if v.Score > 0.08 {
			item.Labels = append(item.Labels, v.Word)
		}
	}

	summary, err := wxyy.GetInstance().Txt2Summary(text)
	if err != nil {
		return
	}
	item.Comment = summary
	return
}

func RegisterToGorse(ctx context.Context, user *dao.Duser) (err error) {
	register := &User{
		UserId:  strconv.FormatInt(user.ID, 10),
		Comment: user.Name,
	}
	defer func(it *User) {
		if _, err = GetInstance().InsertUser(ctx, *it); err != nil {
			log.Println("Insert Err", err)
			return
		}
	}(register)

	return
}

func FavoriteToGorse(ctx context.Context, fa *dao.Dfavorite) (err error) {
	feedback := &Feedback{
		FeedbackType: "star",
		UserId:       strconv.FormatInt(fa.User_id, 10),
		ItemId:       strconv.FormatInt(fa.Video_id, 10),
		Timestamp:    time.Now().Format("2006-01-02"),
	}
	defer func(it *Feedback) {
		if fa.Action_type == "1" {
			feedbacks := []Feedback{*it}
			if _, err = GetInstance().InsertFeedback(ctx, feedbacks); err != nil {
				log.Println("Insert Err", err)
				return
			}
		} else {
			if _, err := GetInstance().DelFeedbackWithUserItem(ctx, feedback.UserId, feedback.ItemId); err != nil {
				log.Println("Del Err", err)
				return
			}
		}
	}(feedback)
	return
}

// func (*GorseClient) CommentToGorse(ctx context.Context, comment *dao.Dcomments) (err error) {
// 	itemPatch := &ItemPatch{}

// 	defer func(it *ItemPatch) {
// 		if _, err = GetInstance().UpdateItem(ctx, strconv.FormatInt(comment.Video_id, 10), *it); err != nil {
// 			log.Println("Update Err", err)
// 			return
// 		}
// 	}(itemPatch)

// 	return
// }

func GetUserRecommend(ctx context.Context, userId int64) (res []int64, err error) {
	if resp, err := GetInstance().GetItemRecommend(ctx, strconv.FormatInt(userId, 10), []string{}, "read", "5s", 5, 0); err != nil {
		log.Println("Get Err", err)
	} else {
		for _, v := range resp {
			id, _ := strconv.ParseInt(v, 10, 64)
			res = append(res, id)
		}
	}
	return
}
