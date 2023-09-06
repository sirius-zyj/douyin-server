package validate

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func InitRateLimit(fillInterval time.Duration, capacity, quantum int64) gin.HandlerFunc {
	limitBucket := ratelimit.NewBucketWithQuantum(fillInterval, capacity, quantum)
	return func(c *gin.Context) {
		if limitBucket.TakeAvailable(1) < 1 {
			c.JSON(http.StatusForbidden, Response{StatusMsg: "detected rate limit"})
			c.Abort()
			return
		}
		c.Next()
	}
}
