package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
//返回给前端的响应格式
{
	"code":10000, //错误码 10000 表示无错误
	"msg":xxx, //提示信息
	"data":{} , // 数据
}
*/

type Resp struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

//RespErr 错误响应
func RespErr(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, Resp{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

//RespErrMsg 自定义错误
func RespErrMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, Resp{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

//RespSuc 成功响应
func RespSuc(c *gin.Context, data interface{}) {
	//CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, Resp{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

//RespNeedBasicAuth 提示进行BasicAuth 认证
func RespNeedBasicAuth(c *gin.Context) {
	c.Header("WWW-Authenticate", "Basic realm=Authorization Required")
	c.AbortWithStatus(http.StatusUnauthorized)
}

//RespJs 响应js文件
func RespJs(c *gin.Context, code string) {
	c.Header("Cache-Control", "nocache")
	c.Header("Pragma", "no-cache")
	c.Data(http.StatusOK, "application/x-javascript", []byte(code))
}
