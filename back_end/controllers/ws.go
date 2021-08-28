package controllers

import (
	"xss/logic"
	"xss/pkg/jwt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//WsHandler 接收Ws连接
func WsHandler(c *gin.Context) {
	//0.接收传参，效验权限
	token := c.Param("jwt")
	mc, ok := jwt.IsValidToken(token, jwt.UserIssuer)
	if !ok {
		zap.L().Error("jwt.IsValidToken failed")
		RespErr(c, CodeServerBusy)
		return
	}
	//1.业务逻辑
	if err := logic.Ws(c, mc.UserID); err != nil {
		zap.L().Error("logic.Ws failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
}
