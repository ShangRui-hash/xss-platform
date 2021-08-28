package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//Conf 存储程序所有的配置信息
var Conf = new(AppConfig)

//AppConfig 配置信息结构体
type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	MachineID    int64  `mapstructure:"machine_id"`
	StartTime    string `mapstructure:"start_time"`
	BaseURL      string `mapstructure:"baseurl"`
	UseHTTPS     bool   `mapstructure:"use_https"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

//LogConfig 日志配置
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

//MySQLConfig mysql配置
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Dbname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

//RedisConfig redis配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	Passwd   string `mapstructure:"passwd"`
	PoolSize int    `mapstructure:"pool_size"`
}

//Init 初始化
func Init(filePath string) error {
	//指定配置文件
	viper.SetConfigFile(filePath)

	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err)
		return err
	}

	//反序列化配置信息
	err = viper.Unmarshal(Conf)
	if err != nil {
		fmt.Printf("viper.Unmarshal(&Conf) failed,err:%v\n", err)
		return err
	}

	//监控配置文件
	viper.WatchConfig()
	//配置文件发生变更之后会调用的回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err := viper.Unmarshal(Conf)
		if err != nil {
			fmt.Printf("viper.Unmarshal(&Conf) failed,err:%v\n", err)
		}
	})
	return nil
}
