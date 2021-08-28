package controllers

import (
	"os"
	"strconv"
	"xss/logic"
	"xss/models"
	"xss/pkg/validate"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//RegisterHandler 注册功能
func RegisterHandler(c *gin.Context) {
	//1.接收参数，后端效验
	var params models.ParamsRegister
	err := ValidateJSONParam(c, &params)
	if err != nil {
		zap.L().Error("user register with invalid param", zap.Error(err))
		return
	}
	//2.效验google验证码
	publicKey := os.Getenv("reCAPTCHA_public_key")
	if len(publicKey) == 0 {
		zap.L().Error("环境变量中未设置谷歌验证码的公钥")
		RespErr(c, CodeServerBusy)
		return
	}
	if !validate.VerifyReCAPTCHAToken(params.ReCAPTCHA, publicKey) {
		zap.L().Error("validate.VerifyReCAPTCHAToken(params.ReCAPTCHA, publicKey) failed")
		RespErr(c, CodeServerBusy)
		return
	}
	//2.业务逻辑
	token, err := logic.Register(params)
	if err != nil {
		zap.L().Error("logic.Register failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	RespSuc(c, token)
}

//UserLoginHandler 用户登录
func UserLoginHandler(c *gin.Context) {
	//1.接收参数
	var params models.ParamLogin
	err := ValidateJSONParam(c, &params)
	if err != nil {
		zap.L().Error("user register with invalid param", zap.Error(err))
		return
	}
	//2.效验google验证码
	publicKey := os.Getenv("reCAPTCHA_public_key")
	if len(publicKey) == 0 {
		zap.L().Error("环境变量中未设置谷歌验证码的公钥")
		RespErr(c, CodeServerBusy)
		return
	}
	if !validate.VerifyReCAPTCHAToken(params.ReCAPTCHA, publicKey) {
		zap.L().Error("validate.VerifyReCAPTCHAToken(params.ReCAPTCHA, publicKey) failed")
		RespErr(c, CodeServerBusy)
		return
	}
	//3.业务逻辑
	token, err := logic.UserLogin(params)
	if err != nil {
		zap.L().Error("logic.UserLogin failed", zap.Error(err))
		RespErr(c, CodeInvalidUserOrPassword)
		return
	}
	RespSuc(c, token)
}

//GetUsersHandler 获取所有用户
func GetUsersHandler(c *gin.Context) {
	//1.接收传参，后端效验
	offset, err := strconv.ParseInt(c.Param("offset"), 10, 64)
	if err != nil {
		RespErr(c, CodeInvalidOffset)
		return
	}
	count, err := strconv.ParseInt(c.Param("count"), 10, 64)
	if err != nil {
		RespErr(c, CodeInvalidCount)
		return
	}
	//2.业务逻辑
	users, err := logic.GetUsers(offset, count)
	if err != nil {
		zap.L().Error("logic.GetUsers failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	RespSuc(c, users)
}

//DeleteUserHandler 删除用户
func DeleteUserHandler(c *gin.Context) {
	//1.接收传参，后端效验
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		RespErr(c, CodeInvalidOffset)
		return
	}
	//2.业务逻辑
	if err := logic.DeleteUser(id); err != nil {
		zap.L().Error("logic.DelteUser failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	RespSuc(c, nil)
}

//SwitchUserStatusHandler 改变用户的状态
func SwitchUserStatusHandler(c *gin.Context) {
	//1.接收传参，后端效验
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		RespErr(c, CodeInvalidOffset)
		return
	}
	//2. 业务逻辑
	if err := logic.SwitchUserStatus(id); err != nil {
		zap.L().Error("logic.BannedUser failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	RespSuc(c, nil)
}
