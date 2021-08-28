package mysql

import (
	"xss/models"

	"go.uber.org/zap"
)

//UserLogin 用户登录
func UserLogin(user *models.User) error {
	sql := `SELECT 
	id,username,password 
	FROM user 
	WHERE username=? AND password = md5(md5(?))`
	err := db.Get(user, sql, user.Username, user.Password)
	if err != nil {
		zap.L().Error("mysql.UserLogin db.Get failed", zap.Error(err))
		return err
	}
	return nil
}

//AddUser 添加用户
func AddUser(user *models.User) error {
	sql := `INSERT INTO user(username,password) VALUES (?,md5(md5(?)))`
	result, err := db.Exec(sql, user.Username, user.Password)
	if err != nil {
		return err
	}
	user.ID, err = result.LastInsertId()
	return err
}

//IsUsernameExist 查询用户名是否存在
func IsUsernameExist(username string) (bool, error) {
	sql := `SELECT count(id) FROM user WHERE username=?`
	var count int
	err := db.Get(&count, sql, username)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

//GetUsers 获取用户
func GetUsers(offset, count int64) ([]models.User, error) {
	sql := `SELECT 
	id,username,created_at,updated_at
	FROM user
	LIMIT ?,?`
	users := make([]models.User, 0, count)
	err := db.Select(&users, sql, offset, count)
	return users, err
}

//GetUser 获取用户详细信息
func GetUser(id int64) (models.User, error) {
	sql := `SELECT * FROM user WHERE id=?`
	var user models.User
	err := db.Get(&user, sql, id)
	return user, err
}

//DeleteUser 删除用户
func DeleteUser(id int64) error {
	sql := `DELETE FROM user WHERE id=?`
	_, err := db.Exec(sql, id)
	return err
}
