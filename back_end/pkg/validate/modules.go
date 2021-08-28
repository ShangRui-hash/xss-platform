package validate

import "xss/dao/mysql"

//IsCommonModuleOrMyModule 是否是公共模块或者是我的模块
func IsCommonModuleOrMyModule(userID, moduleID int64) bool {
	module, err := mysql.GetModule(moduleID)
	if err != nil {
		return false
	}
	if *module.IsCommon == true {
		return true
	}
	return module.UserID == userID
}
