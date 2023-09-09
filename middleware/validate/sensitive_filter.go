package validate

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/importcjj/sensitive"
)

var filter *sensitive.Filter

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func InitSensitive() {
	filter = sensitive.New()
	filter.LoadWordDict("./dict.txt")
	// filter.AddWord("fuck")
	// filter.AddWord("垃圾")
}

func RegisterWordsFilter(c *gin.Context) {
	text := c.Query("username")
	if v, _ := filter.Validate(text); !v {
		c.JSON(http.StatusOK, Response{StatusCode: 404, StatusMsg: "detected seneitive word"})
		c.Abort()
		return
	}
	c.Next()
}

func CommentActionWordsFilter(c *gin.Context) {
	text := c.Query("comment_text")
	if v, _ := filter.Validate(text); !v {
		c.JSON(http.StatusOK, Response{StatusCode: 404, StatusMsg: "detected seneitive word"})
		c.Abort()
		return
	}
	c.Next()
}

func PublishActionWordsFilter(c *gin.Context) {
	text := c.Query("title")
	if v, _ := filter.Validate(text); !v {
		c.JSON(http.StatusOK, Response{StatusCode: 404, StatusMsg: "detected seneitive word"})
		c.Abort()
		return
	}
	c.Next()
}

func MessageActionWordsFilter(c *gin.Context) {
	text := c.Query("content")
	if v, _ := filter.Validate(text); !v {
		c.JSON(http.StatusOK, Response{StatusCode: 404, StatusMsg: "detected seneitive word"})
		c.Abort()
		return
	}
	c.Next()
}
