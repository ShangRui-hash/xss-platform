package middlewares

import (
	"fmt"
	"xss/settings"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"go.uber.org/zap"
)

//TLSHandler  https中间件
func TLSHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     fmt.Sprintf("%s:443", settings.Conf.BaseURL),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			zap.L().Error("secureMiddleware.Process(c.Writer, c.Request)  failed", zap.Error(err))
			c.Abort()
			return
		}
		c.Next()
	}
}
