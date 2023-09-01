package router

import (
	"github.com/gin-gonic/gin"
	"thingue-launcher/common/provider"
)

func BuildRouter() *gin.Engine {
	Router := gin.Default()
	Router.Use(CorsMiddleware())
	//初始化base路由组
	baseGroup := Router.Group(provider.AppConfig.LocalServer.ContentPath)
	{
		//初始化base/ws路由组
		WsRouter.BuildRouter(baseGroup)
		//初始化base/static路由
		StaticRouter.BuildRouter(baseGroup)
	}
	//初始化base/api路由组
	apiGroup := baseGroup.Group("/api")
	{
		//初始化base/api/instance路由组
		InstanceRouter.BuildRouter(apiGroup)
	}
	return Router
}

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
