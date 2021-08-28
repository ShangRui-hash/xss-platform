package controllers

import (
	"xss/logic"
	"xss/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//LogoutHandler 退出登录
func LogoutHandler(c *gin.Context) {
	//1.接收传参,业务逻辑
	var params models.ParamLogout
	err := ValidateJSONParam(c, &params)
	if err != nil {
		zap.L().Error("logout with invalid param", zap.Error(err))
		return
	}
	//2.业务逻辑
	if err := logic.Logout(params.Token); err != nil {
		zap.L().Error("logic.Logout failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	RespSuc(c, nil)
}
