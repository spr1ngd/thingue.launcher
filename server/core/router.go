package core

import (
	"embed"
	"github.com/gin-gonic/gin"
	"thingue-launcher/common/config"
	"thingue-launcher/server/router"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // 缓存预检请求结果的时间，单位为秒
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}

func BuildRouter(staticFiles embed.FS) *gin.Engine {
	apiRouter := router.RouterGroupApp.Api
	wsRouter := router.RouterGroupApp.Ws
	staticRouter := router.RouterGroupApp.Static

	Router := gin.Default()
	Router.Use(CorsMiddleware())
	//初始化base路由组
	baseGroup := Router.Group(config.AppConfig.LocalServer.BasePath)
	{
		//初始化base/ws路由组
		wsRouter.InitWsRouter(baseGroup)
		//初始化base/static路由
		staticRouter.InitStaticRouter(baseGroup, staticFiles)
	}
	//初始化base/api路由组
	apiGroup := baseGroup.Group("/api")
	{
		//初始化base/api/agent路由组
		apiRouter.InitAgentRouter(apiGroup)
		//初始化base/api/admin路由组
		apiRouter.InitAdminRouter(apiGroup)
	}
	return Router
}
