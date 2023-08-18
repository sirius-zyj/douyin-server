package main

import (
	"context"
	"douyin-server/dao"
	comment "douyin-server/rpc/kitex_gen/comment"
	user "douyin-server/rpc/kitex_gen/user"
	"log"
	"strconv"
	"strings"
	"time"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	resp = new(comment.DouyinCommentActionResponse)

	token, videoId, actionType, commentText, commentId := req.Token, req.VideoId, req.ActionType, *req.CommentText, *req.CommentId

	index := strings.Index(token, "*")
	user_id, _ := strconv.ParseInt(token[index+1:], 10, 64)
	//最好有一个用户身份验证过程
	if actionType == 1 {
		dComment := dao.Dcomments{
			Comment_text: commentText,
			User_id:      user_id,
			Video_id:     videoId,
			Created_at:   time.Now(),
		}

		err = dao.InsertComment(dComment)
		if err != nil {
			log.Println("插入评论失败")
			resp.StatusCode = 404
			resp.StatusMsg = new(string)
			*resp.StatusMsg = "插入评论失败"
			return
		}
		User, _ := dao.GetUserById(user_id)

		resp.StatusCode = 0
		resp.StatusMsg = new(string)
		*resp.StatusMsg = "插入评论成功"
		resp.Comment = &comment.Comment{
			Id: dComment.Id,
			User: &user.User{
				Id:   User.Id,
				Name: User.Name,
			},
			Content:    dComment.Comment_text,
			CreateDate: dComment.Created_at.Format("01-02"),
		}
	} else {
		err := dao.EraseComment(commentId)
		if err != nil {
			log.Println("删除评论失败")
			resp.StatusCode = 404
			resp.StatusMsg = new(string)
			*resp.StatusMsg = "删除评论失败"
		} else {
			log.Println("删除评论成功")
			resp.StatusCode = 0
			resp.StatusMsg = new(string)
			*resp.StatusMsg = "删除评论成功"
		}
	}
	return
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	resp = new(comment.DouyinCommentListResponse)

	videoId := req.VideoId
	CommentList, err := dao.GetAllComments(videoId)
	if err != nil {
		log.Println("拉取评论失败")
		resp.StatusCode = 404
		resp.StatusMsg = new(string)
		*resp.StatusMsg = "拉取评论失败,请重试"
	} else {
		log.Println("拉取评论成功")
		for _, tmp := range CommentList {
			//---------获取评论的用户-------------
			temp_user, _ := dao.GetUserById(tmp.User_id)
			resp.CommentList = append(resp.CommentList, &comment.Comment{
				Id: tmp.Id,
				User: &user.User{
					Id:   temp_user.Id,
					Name: temp_user.Name,
				},
				Content:    tmp.Comment_text,
				CreateDate: tmp.Created_at.Format("01-02"),
			})
		}
		resp.StatusCode = 0
		resp.StatusMsg = new(string)
		*resp.StatusMsg = "拉取评论成功"
	}
	return
}
