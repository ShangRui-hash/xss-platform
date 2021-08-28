package mysql

import (
	"errors"
	"fmt"
	"os"
	"time"
	"xss/settings"

	_ "github.com/go-sql-driver/mysql" //匿名导入 默认会自动执行该包中的init()方法
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

//Init 初始化
func Init(cfg *settings.MySQLConfig) (err error) {
	username := os.Getenv("mysql_username")
	if len(username) == 0 {
		zap.L().Error("请在环境变量中设置mysql的username,export mysql_username=你的mysql用户名")
		return errors.New("mysql_username is null")
	}
	password := os.Getenv("mysql_password")
	if len(password) == 0 {
		zap.L().Error("请在环境变量中设置mysql的password,export mysql_password=你的密码")
		return errors.New("mysql_password is null")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		username,
		password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname,
	)

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return err
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	db.SetConnMaxLifetime(time.Duration(viper.GetInt("mysql.max_left_time")) * time.Second)
	return nil
}

//Close 关闭
func Close() {
	db.Close()
}
