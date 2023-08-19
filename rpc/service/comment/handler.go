package main

import (
	"context"
	"douyin-server/dao"
	comment "douyin-server/rpc/kitex_gen/comment"
	"log"
	"strconv"
	"strings"
	"time"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

func setCommentActionResp(resp *comment.DouyinCommentActionResponse, statusCode int32, statusMsg string) {
	resp.StatusCode = statusCode
	resp.StatusMsg = new(string)
	*resp.StatusMsg = statusMsg
}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	resp = new(comment.DouyinCommentActionResponse)

	token, videoId, actionType := req.Token, req.VideoId, req.ActionType

	index := strings.Index(token, "*")
	user_id, _ := strconv.ParseInt(token[index+1:], 10, 64)
	//最好有一个用户身份验证过程 TODO
	if actionType == 1 {
		dComment := dao.Dcomments{
			Comment_text: *req.CommentText,
			User_id:      user_id,
			Video_id:     videoId,
			Created_at:   time.Now(),
		}

		err = dao.InsertComment(dComment)
		if err != nil {
			log.Println("插入评论失败")
			setCommentActionResp(resp, 404, "插入评论失败")
			return
		}
		dao.UpdateFeed("id", videoId, "comment_count", 1) //评论数+1
		tmp, _ := dao.GetUserById(user_id)
		resp.Comment = &comment.Comment{
			Id:         dComment.Id,
			User:       dao.DaoUser2RPCUser(&tmp),
			Content:    dComment.Comment_text,
			CreateDate: dComment.Created_at.Format("01-02"),
		}
		setCommentActionResp(resp, 0, "插入评论成功")
	} else {
		commentId, _ := strconv.ParseInt(*req.CommentId, 10, 64)
		err := dao.EraseComment(commentId, videoId)
		if err != nil {
			log.Println("删除评论失败")
			setCommentActionResp(resp, 404, "删除评论失败")
		} else {
			dao.UpdateFeed("id", videoId, "comment_count", -1) //评论数-1
			log.Println("删除评论成功")
			setCommentActionResp(resp, 0, "删除评论成功")
		}
	}
	return
}

func setCommentListResp(resp *comment.DouyinCommentListResponse, statusCode int32, statusMsg string) {
	resp.StatusCode = statusCode
	resp.StatusMsg = new(string)
	*resp.StatusMsg = statusMsg
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	resp = new(comment.DouyinCommentListResponse)

	videoId := req.VideoId
	CommentList, err := dao.GetAllComments(videoId)
	if err != nil {
		log.Println("拉取评论失败")
		setCommentListResp(resp, 404, "拉取评论失败,请重试")
	} else {
		log.Println("拉取评论成功")
		for _, tmp := range CommentList {
			//---------获取评论的用户-------------
			temp_user, _ := dao.GetUserById(tmp.User_id)
			resp.CommentList = append(resp.CommentList, &comment.Comment{
				Id:         tmp.Id,
				User:       dao.DaoUser2RPCUser(&temp_user),
				Content:    tmp.Comment_text,
				CreateDate: tmp.Created_at.Format("01-02"),
			})
		}
		setCommentListResp(resp, 0, "拉取评论成功")
	}
	return
}
