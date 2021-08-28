package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"xss/controllers"
	"xss/dao/mysql"
	"xss/dao/redis"
	"xss/logger"
	"xss/pkg/snowflake"
	"xss/routes"
	"xss/settings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @title XSS Platform
// @version 1.0
// @description 基于gin+vue+mysql+redis 实现的xss平台
// @termsOfService http://swagger.io/terms/

// @contact.name Rick Shang
// @contact.url https://gitee.com/nothing-is-nothing
// @contact.email 2227627947@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath
func main() {
	var defaultConfigFile string
	GinMode := os.Getenv(gin.EnvGinMode)
	if GinMode == gin.ReleaseMode {
		defaultConfigFile = "./conf/config.production.json"
	} else {
		defaultConfigFile = "./conf/config.development.json"
	}
	//0.接收命令行参数
	var configFilePath string
	flag.StringVar(&configFilePath, "config", defaultConfigFile, "指定configPath")
	flag.Parse()

	//1.加载配置
	if err := settings.Init(configFilePath); err != nil {
		fmt.Printf("init setting failed,err:%v\n", err)
		return
	}

	//2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed,err:%v\n", err)
		return
	}
	defer zap.L().Sync() //同步到日志

	//3.初始化MySQL连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed,err:%v\n", err)
		return
	}
	defer mysql.Close()

	//4.初始化Redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed,err:%v\n", err)
		return
	}
	defer redis.Close()
	//5.初始化分布式ID生成器
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed,err:%v\n", err)
		return
	}
	//6.初始化gin框架 binding validate 使用的翻译器
	if err := controllers.InitTrans("zh"); err != nil {
		fmt.Printf("init translator failed,err:%v\n", err)
		return
	}
	//8.注册路由
	router := routes.Setup(settings.Conf.Mode)

	//9.启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: router,
	}
	var httpsSrv *http.Server
	if settings.Conf.UseHTTPS {
		httpsSrv = &http.Server{
			Addr:    ":443",
			Handler: router,
		}
	}

	// 开启一个goroutine启动服务
	go func() {
		go func() {
			err := srv.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		if settings.Conf.UseHTTPS {
			certFile := os.Getenv("cert_file")
			if 0 == len(certFile) {
				certFile = "./SSL/test.pem"
			}
			keyFile := os.Getenv("key_file")
			if 0 == len(keyFile) {
				keyFile = "./SSL/test.key"
			}
			err := httpsSrv.ListenAndServeTLS(certFile, keyFile)
			if err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}

	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")

	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("HTTP Server Shutdown: ", zap.Error(err))
	}
	if settings.Conf.UseHTTPS {
		if err := httpsSrv.Shutdown(ctx); err != nil {
			zap.L().Fatal("HTTPS Server Shutdown: ", zap.Error(err))
		}
	}

	zap.L().Info("Server exiting")
}
