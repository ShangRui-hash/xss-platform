package logic

import (
	"encoding/json"
	"strconv"
	"xss/dao/mysql"
	"xss/dao/redis"
	"xss/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//GetLoots 获取项目所有的战利品
func GetLoots(URLKey string) ([]models.Loot, error) {
	return mysql.GetLoots(URLKey)
}

//DelMyLoot 删除我的战利品
func DelMyLoot(id int64) (err error) {
	return mysql.DelMyLoot(id)
}

//RecvLoot 接收xss payload 传来的数据
func RecvLoot(c *gin.Context, URLKey string) error {
	//0.查询该项目属于哪个用户
	userID, err := mysql.GetOwnerOfProject(URLKey)
	if err != nil {
		zap.L().Error("mysql.GetOwnerOfProject failed", zap.Error(err))
		return err
	}
	//1.查询本项目的所有模块
	moduleIDs, err := redis.NewModuleSetOfProject(URLKey).Members()
	if err != nil {
		zap.L().Error("redis.NewModuleSetOfProject failed", zap.Error(err))
		return err
	}
	//2.查询这些模块都接收什么参数
	allParams := make([]string, 0)
	for _, moduleID := range moduleIDs {
		id, err := strconv.ParseInt(moduleID, 10, 64)
		if err != nil {
			zap.L().Error("strconv.ParseInt failed", zap.Error(err))
			return err
		}
		params, err := redis.NewParamSet(id).Members()
		if err != nil {
			zap.L().Error("redis.NewParamSet(id).Members failed", zap.Error(err))
			return err
		}
		allParams = append(allParams, params...)
	}
	//3.接收这些参数的值并存储到数据库中
	paramMap := make(map[string]string)
	var ok bool
	for _, param := range allParams {
		paramMap[param], ok = c.GetQuery(param)
		if !ok {
			paramMap[param] = c.PostForm(param)
		}
	}
	lootContent, err := json.Marshal(paramMap)
	if err != nil {
		zap.L().Error("json.Marshal failed", zap.Error(err))
		return err
	}
	Loot := models.Loot{
		URLKey:  URLKey,
		Content: string(lootContent),
	}
	if err := mysql.AddLoot(&Loot); err != nil {
		zap.L().Error("mysql.AddLoot failed", zap.Error(err))
		return err
	}
	//4.送入战利品管道,以便websocket服务端向客户端进行推送
	lootStr, err := json.Marshal(Loot)
	if err != nil {
		zap.L().Error("json.Marshal failed", zap.Error(err))
		return err
	}
	if err := redis.NewLootQueue(userID).EnQueue(string(lootStr)); err != nil {
		zap.L().Error("redis.NewLootQueue.EnQueue failed", zap.Error(err))
		return err
	}
	return nil
}
