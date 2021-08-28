package logic

import (
	"errors"
	"xss/dao/mysql"
	"xss/dao/redis"
	"xss/models"
	"xss/pkg/jwt"
	"xss/pkg/randstr"

	"go.uber.org/zap"
)

//Register 用户注册
func Register(params models.ParamsRegister) (token string, err error) {
	//0.生成用户名
	username, err := GenUsername()
	if err != nil {
		zap.L().Error("GenUsername failed", zap.Error(err))
		return "", err
	}
	//1.信息入库
	user := models.User{
		Username: username,
		Password: params.Password,
	}
	if err := mysql.AddUser(&user); err != nil {
		zap.L().Error("mysql.AddUser failde", zap.Error(err))
		return "", err
	}
	//2.更新登录时间
	_, err = redis.NewLoginTime().UpdateLoginTime(user.Username)
	if err != nil {
		zap.L().Error("redis.NewLoginTime().UpdateLoginTime failed", zap.Error(err))
		return "", err
	}
	//3.颁发令牌
	return jwt.GenUserToken(user.ID, user.Username)
}

//GenUsername 生成一个数据库中没有的用户名
func GenUsername() (username string, err error) {
	for {
		username := randstr.String(10)
		exist, err := mysql.IsUsernameExist(username)
		if err != nil {
			return "", err
		}
		if exist == false {
			return username, nil
		}
	}
}

//UserLogin 用户登录
func UserLogin(params models.ParamLogin) (token string, err error) {
	//1.效验用户名的错误次数
	counter := redis.NewCounter(params.Username)
	val, err := counter.Get()
	if err != nil {
		zap.L().Error("counter.Get failed", zap.Error(err))
		return "", err
	}
	if val >= 10 {
		return "", errors.New("用户名密码错误次数30分钟内已超过10次")
	}
	//2.效验用户名是否存在
	exist, err := mysql.IsUsernameExist(params.Username)
	if err != nil {
		return "", err
	}
	if !exist {
		return "", errors.New("用户名不存在")
	}
	//3.查询密码是否正确
	user := models.User{
		Username: params.Username,
		Password: params.Password,
	}
	if err = mysql.UserLogin(&user); err != nil {
		zap.L().Error("mysql.UserLogin failed", zap.Error(err))
		counter.Incr() //增加错误次数
		return "", err
	}
	//4.更新登录时间
	_, err = redis.NewLoginTime().UpdateLoginTime(user.Username)
	if err != nil {
		zap.L().Error("redis.NewLoginTime().UpdateLoginTime failed", zap.Error(err))
		return "", err
	}
	//5.生成token
	return jwt.GenUserToken(user.ID, user.Username)
}

//GetUsers 获取用户
func GetUsers(offset, count int64) ([]models.User, error) {
	users, err := mysql.GetUsers(offset, count)
	if err != nil {
		return users, err
	}
	loginTime := redis.NewLoginTime()
	blacklist := redis.NewUserBlackList()
	for i := range users {
		//获取登录时间
		users[i].LoginedAt, err = loginTime.Get(users[i].Username)
		if err != nil {
			zap.L().Error("loginTime.Get failed", zap.Error(err))
		}
		//获取是否被封号
		users[i].IsBanned, err = blacklist.IsMember(users[i].Username)
		if err != nil {
			zap.L().Error("blacklist.IsMember failed", zap.Error(err))
		}
	}
	return users, err
}

//DeleteUser 删除用户
func DeleteUser(id int64) error {
	return mysql.DeleteUser(id)
}

//SwitchUserStatus 改变用户状态
func SwitchUserStatus(id int64) error {
	user, err := mysql.GetUser(id)
	if err != nil {
		zap.L().Error("mysql.GetUser failed", zap.Error(err))
		return err
	}
	blackList := redis.NewUserBlackList()
	exist, err := blackList.IsMember(user.Username)
	if err != nil {
		zap.L().Error("blackList.IsMember(user.Username) failed", zap.Error(err))
		return err
	}
	if exist {
		_, err = blackList.Rem(user.Username)
		if err != nil {
			zap.L().Error("blackList.Rem failed", zap.Error(err))
			return err
		}
	} else {
		_, err = blackList.Add(user.Username)
		if err != nil {
			zap.L().Error("blackList.Add failed", zap.Error(err))
			return err
		}
	}
	return nil
}
