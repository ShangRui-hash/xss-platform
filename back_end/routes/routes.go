package routes

import (
	"fmt"
	"net/http"
	"time"
	"xss/controllers"
	"xss/logger"
	"xss/middlewares"
	"xss/settings"

	_ "xss/docs" // 千万不要忘了导入把你上一步生成的docs

	"github.com/gin-gonic/gin"
)

//Setup 配置路由
func Setup(mode string) *gin.Engine {
	//一共三种模式，其余2种模式都当做调试模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //gin设置为发布模式,就不会输出Gin-debug信息了
	}
	r := gin.New()
	fmt.Println("settings.Conf.UseHTTPS", settings.Conf.UseHTTPS)
	if true == settings.Conf.UseHTTPS {
		r.Use(middlewares.TLSHandler())
	}
	//日志中间件、限流中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(time.Millisecond, 1000))
	//前端
	r.Static("/js", "../front_end/dist/js")
	r.Static("/css", "../front_end/dist/css")
	r.Static("/fonts", "../front_end/dist/fonts")
	//解析模板
	r.LoadHTMLGlob("../front_end/dist/index.html")
	//渲染模板
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//后台API
	admin := r.Group("/api/v1/admin")
	//管理员登录
	admin.POST("/login", controllers.AdminLoginHandler)
	//json web token 效验中间件
	admin.Use(middlewares.JWTAuthAdminMiddleware())
	{
		//添加模块
		admin.POST("/modules", controllers.AddModuleHandler)
		//获取模块列表
		admin.GET("/modules/:offset/:count", controllers.GetModulesHandler)
		//获取模块详情
		admin.GET("/module/:id", controllers.AdminGetModuleDetailHandler)
		//更新模块
		admin.PUT("/module", controllers.UpdateModuleHandler)
		//删除模块
		admin.DELETE("/module/:id", controllers.DeleteModuleHandler)
		//查看用户
		admin.GET("/users/:offset/:count", controllers.GetUsersHandler)
		//删除用户
		admin.DELETE("/user/:id", controllers.DeleteUserHandler)
		//改变用户状态
		admin.PUT("/user/:id", controllers.SwitchUserStatusHandler)
	}

	//前台API
	home := r.Group("api/v1/")
	//用户注册
	home.POST("/register", controllers.RegisterHandler)
	//用户登录
	home.POST("/login", controllers.UserLoginHandler)
	//退出登录 (前后台共用)
	home.POST("/logout", controllers.LogoutHandler)
	//websocket
	home.GET("/ws/:jwt", controllers.WsHandler)
	//获取添加模块的注意事项
	home.GET("/module/notice", controllers.GetModuleNoticeHandler)
	//jwt认证中间件
	home.Use(middlewares.JWTAuthUserMiddleware())
	{
		//获取公共模块列表
		home.GET("/modules/common/:offset/:count", controllers.GetCommonModulesHandler)
		//获取某个公共模块 或 我的模块的详情
		home.GET("/module/:id", controllers.UserGetModuleDetailHandler)
		//添加我的模块
		home.POST("/modules", controllers.AddModuleHandler)
		//查看我的模块
		home.GET("/modules/:offset/:count", controllers.GetMyModulesHandler)
		//删除我的模块
		home.DELETE("/module/:id", controllers.DeleteMyModuleHandler)
		//修改我的模块
		home.PUT("/module", controllers.UpdateMyModuleHandler)
		//获取创建项目的表单
		home.GET("/projectform", controllers.GetProjectFormHandler)
		//创建项目
		home.POST("/project", controllers.CreateProjectHandler)
		//获取项目列表
		home.GET("/projects/:offset/:count", controllers.GetProjectListHandler)
		//删除项目
		home.DELETE("/project/:id", controllers.DeleteProjectHandler)
		//修改项目
		home.PUT("/project/:url_key", controllers.UpdateProjectHandler)
		//获取项目的战利品
		home.GET("/loots/:url_key", controllers.GetLootsHandler)
		//删除战利品
		home.DELETE("/loot/:id", controllers.DelMyLootHandler)
	}
	//攻击模块
	attack := r.Group("/")
	{
		//获取对应项目的xss payload
		attack.GET("/:url_key", controllers.ProvideXSSPayloadHandler)
		//BasicAuth 钓鱼
		attack.GET("/basicAuth", controllers.BasicAuthPhishHandler)
		//接收GET传参的数据
		attack.GET("/loot", controllers.RecvLootHandler)
		//接收POST传参的数据
		attack.POST("/postloot", controllers.RecvLootHandler)
		attack.GET("/evil", controllers.EvilHandler)

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	//生产环境下记得关闭
	// r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	return r
}
