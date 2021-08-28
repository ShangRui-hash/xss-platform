package redis

import (
	"fmt"

	"github.com/spf13/viper"
)

//OptionOfProject 项目和模块参数的关联表
type OptionOfProject struct {
	key    string
	delKey string
}

//NewOptionOfProject 关联项目和模块参数
func NewOptionOfProject(projectURLKey string, moduleID int64) *OptionOfProject {
	return &OptionOfProject{
		key: fmt.Sprintf("%s::project::%s::moduleID::%d", viper.GetString("redis.key_prefix"), projectURLKey, moduleID),
	}
}

//Set 设置配置项的值
func (o *OptionOfProject) Set(name string, value string) (bool, error) {
	return rdb.HSet(o.key, name, value).Result()
}

//Get 获取配置项的值
func (o *OptionOfProject) Get(name string) (string, error) {
	return rdb.HGet(o.key, name).Result()
}

//GetAll 获取所有配置项的值
func (o *OptionOfProject) GetAll() (map[string]string, error) {
	return rdb.HGetAll(o.key).Result()
}

//Clear 清除
func (o *OptionOfProject) Clear() error {
	_, err := rdb.Del(o.key).Result()
	return err
}

//ClearAllOptions 删除掉该项目相关的所有参数
func ClearAllOptions(projectURLKey string) error {
	delKey := fmt.Sprintf("%s::project::%s::moduleID::*", viper.GetString("redis.key_prefix"), projectURLKey)
	_, err := rdb.Del(delKey).Result()
	return err
}
