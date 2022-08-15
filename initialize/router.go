package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/miaojiaxi2004/go_server/global"
	"github.com/miaojiaxi2004/go_server/middleware"
	"github.com/miaojiaxi2004/go_server/model/system"
	"github.com/miaojiaxi2004/go_server/router"
)

// 初始化总路由(仅供测试，到时候放到model目录下)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	PublicGroup := Router.Group("")
	global.LOG.Info("use middleware logger")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	//Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	global.LOG.Info("use middleware cors")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			var test []system.CommunityIndex
			global.DB.Find(&test, "status = ?", "1")
			c.JSON(200, test)
		})
	}
	{
		//systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		//systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	global.LOG.Info("register swagger handler")
	{
		systemRouter.InitApiRouter(PrivateGroup) // 注册功能api路由

		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由

	}
	global.LOG.Info("router register success")
	return Router
}
