package controllers

import (
	"strconv"
	"xss/dao/mysql"
	"xss/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//GetLootsHandler 获取项目的战利品
func GetLootsHandler(c *gin.Context) {
	//1.接收参数
	URLKey := c.Param("url_key")
	if len(URLKey) != 4 {
		RespErr(c, CodeInvalidParam)
		return
	}
	//2.业务逻辑
	loots, err := logic.GetLoots(URLKey)
	if err != nil {
		zap.L().Error("logic.GetLoots failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}

	RespSuc(c, loots)
}

//RecvLootHandler 接收战利品
func RecvLootHandler(c *gin.Context) {
	//1.接收参数
	URLKey := c.Query("url_key")
	if len(URLKey) != 4 {
		RespErr(c, CodeServerBusy)
		return
	}
	//2.业务逻辑
	err := logic.RecvLoot(c, URLKey)
	if err != nil {
		zap.L().Error("logic.RecvData failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	RespSuc(c, nil)
}

//DelMyLootHandler 删除我的战利品
func DelMyLootHandler(c *gin.Context) {
	//1.接收参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		RespErr(c, CodeServerBusy)
		return
	}
	userID, exist := c.Get(CtxUserIDKey)
	if !exist {
		zap.L().Error("userid don't exist")
		RespErr(c, CodeServerBusy)
		return
	}
	//2.后端效验
	if false == mysql.ValidateLootOwner(userID.(int64), id) {
		zap.L().Error("false == mysql.ValidateLootOwner(userID, id) ")
		RespErr(c, CodeServerBusy)
		return
	}
	//3.业务逻辑
	if err := logic.DelMyLoot(id); err != nil {
		zap.L().Error("logic.DelMyLoot failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	RespSuc(c, nil)
}
