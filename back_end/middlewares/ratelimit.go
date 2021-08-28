package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

//RateLimitMiddleware 令牌桶限流
//@fillInterval 向令牌桶中放令牌的速率
//@cap 令牌桶的总容量，令牌桶一开始是满的
func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		//如果取不到令牌
		if bucket.TakeAvailable(1) == 0 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		//取到令牌放行
		c.Next()
	}
}
