package validate

import (
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 判断文件扩展名是否为视频文件
func isVideoFile(file *multipart.FileHeader) bool {
	ext := strings.ToLower(file.Filename[strings.LastIndex(file.Filename, ".")+1:])
	return ext == "mp4" || ext == "avi" || ext == "mkv" // 根据需要添加其他视频文件扩展名
}

func VideoFilter(c *gin.Context) {
	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 404, StatusMsg: "参数错误"})
		c.Abort()
		return
	}
	// 限制文件大小
	if file.Size > 20*1024*1024 { // 假设限制文件大小为 10MB
		c.JSON(http.StatusBadRequest, Response{StatusCode: 404, StatusMsg: "文件过大"})
		c.Abort()
		return
	}
	// 限制文件类型
	// 判断文件扩展名是否为视频文件
	if !isVideoFile(file) {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 404, StatusMsg: "不是视频文件"})
		c.Abort()
		return
	}
}
