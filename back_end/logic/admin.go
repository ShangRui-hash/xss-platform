package logic

import (
	"xss/models"
	"xss/pkg/jwt"
	"xss/dao/mysql"
	"go.uber.org/zap"
)

//AdminLogin 管理员登录的业务处理函数
func AdminLogin(p *models.ParamLogin) (string, error) {
	//查询用户名是否存在
	exist, err := mysql.IsAdminExist(p.Username)
	if err != nil {
		zap.L().Error("mysql.IsAdminExist(p.Username) failed", zap.Error(err))
		return "", err
	}
	if false == exist {
		zap.L().Error("mysql.IsAdminExist(p.Username)", zap.Error(mysql.ErrAdminNotExist))
		return "", mysql.ErrAdminNotExist
	}
	//查询密码是否正确
	admin := models.Admin{
		Username: p.Username,
		Password: p.Password,
	}
	err = mysql.AdminLogin(&admin)
	if err != nil {
		return "", err
	}
	//颁发令牌
	return jwt.GenAdminToken(admin.ID, admin.Username)
}
