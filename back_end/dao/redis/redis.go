package redis

import (
	"fmt"
	"xss/settings"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// Init 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Passwd, // no password set
		DB:       cfg.Db,     // use default DB
		PoolSize: cfg.PoolSize,
	})

	_, err = rdb.Ping().Result()
	return err
}

//Close 关闭连接
func Close() {
	rdb.Close()
}
