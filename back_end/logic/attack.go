package logic

import (
	"fmt"
	"strconv"
	"strings"
	"xss/dao/mysql"
	"xss/dao/redis"

	"github.com/spf13/viper"
)

//ProvideXSSPayload 提供项目对应的xss payload
func ProvideXSSPayload(URLKey string) (payload string, err error) {
	//1.获取该项目关联的所有模块
	modulesIDs, err := redis.NewModuleSetOfProject(URLKey).Members()
	if err != nil {
		return "", err
	}

	//2.将所有模块的payload拼接起来
	for _, moduleID := range modulesIDs {
		id, err := strconv.ParseInt(moduleID, 10, 64)
		if err != nil {
			return "", err
		}
		module, err := mysql.GetModule(id)
		if err != nil {
			return "", err
		}
		//获取该模块所有配置项的值
		options, err := redis.NewOptionOfProject(URLKey, module.ID).GetAll()
		if err != nil {
			return "", err
		}
		//替换掉payload中的配置项
		for key, val := range options {
			module.XSSPayload = strings.ReplaceAll(module.XSSPayload, fmt.Sprintf("{option_%s}", key), val)
		}
		//替换掉payload中的base_url
		module.XSSPayload = strings.ReplaceAll(module.XSSPayload, viper.GetString("payload.baseurl"), viper.GetString("baseurl"))
		//替换掉payload中的url_key
		module.XSSPayload = strings.ReplaceAll(module.XSSPayload, viper.GetString("payload.url_key"), URLKey)
		payload = payload + "\n" + module.XSSPayload
	}
	return payload, nil
}
