package logic

import "xss/dao/redis"

//Logout 退出登录
func Logout(token string) error {
	_, err := redis.NewTokenBlackList().Add(token)
	return err
}
