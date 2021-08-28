package mysql

import (
	"fmt"
	"xss/models"

	"go.uber.org/zap"
)

//GetCommonModules 获取公共模块
func GetCommonModules(offset, count int64) ([]models.Module, error) {
	modules := make([]models.Module, 0)
	sql := `select id,name,description,created_at,updated_at from modules where is_common=1 limit ?,?`
	err := db.Select(&modules, sql, offset, count)
	if err != nil {
		zap.L().Error("db.Get(&modules, sql, offset, count) failed", zap.Error(err))
		return nil, err
	}
	return modules, nil
}

//GetMyModules 获取我的模块
func GetMyModules(offset, count, userID int64) ([]models.Module, error) {
	modules := make([]models.Module, 0)
	sql := `select id,name,description,created_at,updated_at,is_common from modules where user_id=? and user_type=0 limit ?,?`
	err := db.Select(&modules, sql, userID, offset, count)
	if err != nil {
		zap.L().Error("db.Get(&modules, sql, offset, count) failed", zap.Error(err))
		return nil, err
	}
	return modules, nil
}

//GetAllModules 获取指定用户可以查看的所有模块（公共模块+自己的模块）
func GetAllModules(userid int64) ([]models.Module, error) {
	modules := make([]models.Module, 0)
	sql := `select id,name,description,created_at,updated_at from modules where is_common=1 or (user_id = ? and user_type=0)`
	err := db.Select(&modules, sql, userid)
	if err != nil {
		zap.L().Error("db.Get(&modules, sql, offset, count) failed", zap.Error(err))
		return nil, err
	}
	return modules, nil
}

//GetModules 获取模块
func GetModules(offset, count int64) ([]models.Module, error) {
	modules := make([]models.Module, 0)
	sql := `select id,name,description,user_id,user_type,created_at,updated_at,is_common from modules limit ?,?`
	err := db.Select(&modules, sql, offset, count)
	if err != nil {
		zap.L().Error("db.Get(&modules, sql, offset, count) failed", zap.Error(err))
		return nil, err
	}
	tableName := map[bool]string{
		true:  "admin",
		false: "user",
	}
	var username string
	for i := range modules {
		sql = fmt.Sprintf(`select username from %s where id=?`, tableName[*modules[i].IsAdmin])
		err := db.Get(&username, sql, modules[i].UserID)
		if err != nil {
			zap.L().Error("db.Get(&username, sql, tableName[modules[i].IsAdmin], modules[i].ID) failed", zap.Error(err))
			return nil, err
		}
		modules[i].Username = username
	}
	return modules, nil
}

//AddModule 添加模块
func AddModule(m *models.Module) (err error) {
	sql := `insert into 
	modules(name,description,xss_payload,user_id,user_type,is_common) 
	values(?,?,?,?,?,?) `
	result, err := db.Exec(sql, m.Name, m.Desc, m.XSSPayload, m.UserID, m.IsAdmin, m.IsCommon)
	if err != nil {
		return err
	}
	m.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

//GetModule 获取模块详情
func GetModule(id int64) (models.Module, error) {
	var m models.Module
	sql := `select * from modules where id=?`
	err := db.Get(&m, sql, id)
	if err != nil {
		return m, err
	}
	return m, nil
}

//UpdateModule 修改模块
func UpdateModule(param models.ParamsUpdateModule) error {
	sql := `update modules set name=?,description=?,xss_payload=?,is_common=? where id=?`
	_, err := db.Exec(sql, param.Name, param.Desc, param.XSSPayload, param.IsCommon, param.ID)
	return err
}

//DeleteModule 删除模块
func DeleteModule(id int64) error {
	sql := `delete from modules where id=?`
	_, err := db.Exec(sql, id)
	return err
}

//DeleteMyModule 删除指定用户的模块
func DeleteMyModule(moduleID, userID int64) error {
	sql := `delete from modules where id=? and user_id =? and user_type=0`
	_, err := db.Exec(sql, moduleID, userID)
	return err
}

//ValidateModuleOwner 效验模块的拥有者
func ValidateModuleOwner(userID, moduleID int64) bool {
	module, err := GetModule(moduleID)
	if err != nil {
		return false
	}
	return module.UserID == userID
}
