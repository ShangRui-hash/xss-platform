package mysql

import (
	"xss/models"

	"go.uber.org/zap"
)

//IsAdminExist 判断管理员账户是否存在
func IsAdminExist(phone string) (bool, error) {
	sql := `select count(id) from admin where username= ?`
	var count int
	err := db.Get(&count, sql, phone)
	if err != nil {
		zap.L().Error("mysql.IsAdminExist db.Get() failed", zap.Error(err))
		return false, err
	}
	if count == 1 {
		return true, nil
	}
	return false, nil
}

//AdminLogin 管理员登录
func AdminLogin(admin *models.Admin) error {
	sql := `select id,username,password from admin where username=? and password = md5(md5(?))`
	err := db.Get(admin, sql, admin.Username, admin.Password)
	if err != nil {
		zap.L().Error("mysql.AdminLogin db.Get failed", zap.Error(err))
		return err
	}
	return nil
}
