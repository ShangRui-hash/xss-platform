package controllers

import (
	"os"
	"xss/logic"
	"xss/models"
	"xss/pkg/validate"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AdminLoginHandler 登录
// @Summary 管理员登录接口
// @Description 管理员登录接口
// @Tags 后台相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamLogin true "查询参数"
// @Success 200 {object} _RespAdminLogin
// @Router /api/v1/admin/login [post]
func AdminLoginHandler(c *gin.Context) {
	//1.接受参数,后端效验
	var params models.ParamLogin
	err := ValidateJSONParam(c, &params)
	if err != nil {
		zap.L().Error("Admin login with invalid param", zap.Error(err))
		return
	}
	//2.效验谷歌验证码
	publicKey := os.Getenv("reCAPTCHA_public_key")
	if len(publicKey) == 0 {
		zap.L().Error("环境变量中未设置谷歌验证码的公钥")
		RespErr(c, CodeServerBusy)
		return
	}
	if !validate.VerifyReCAPTCHAToken(params.ReCAPTCHA, publicKey) {
		zap.L().Error("validate.VerifyReCAPTCHAToken(params.ReCAPTCHA, publicKey) == fasle")
		RespErr(c, CodeServerBusy)
		return
	}
	//2.业务逻辑部分
	token, err := logic.AdminLogin(&params)
	if err != nil {
		zap.L().Error("logic.AdminLogin failed", zap.Error(err))
		RespErr(c, CodeLoginFailed)
		return
	}
	RespSuc(c, token)
}
