package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/provider"
	"thingue-launcher/server/core/router"
)

var httpHandler http.Handler

func CreateGinRouter() {
	Router := gin.Default()
	Router.Use(corsMiddleware())
	//初始化base路由
	baseGroup := Router.Group(provider.AppConfig.LocalServer.ContentPath)
	{
		//初始化base/grpc路由组
		baseGroup.Any("/grpc-gateway/*path", gin.WrapH(createGrpcGatewayHandler()))
		//初始化base/ws路由组
		router.WsRouter.BuildRouter(baseGroup)
		//初始化base/static路由
		buildStaticRouter(baseGroup)
		//初始化base/api路由
		apiGroup := baseGroup.Group("/api")
		{
			//初始化base/api/*路由组
			router.InstanceRouter.BuildRouter(apiGroup)
			router.SyncRouter.BuildRouter(apiGroup)
		}
	}
	httpHandler = Router
}

func buildStaticRouter(group *gin.RouterGroup) {
	group.Static("/storage", "./thingue-launcher/storage")
	if provider.AppConfig.LocalServer.UseExternalStatic {
		group.Static("/static", provider.AppConfig.LocalServer.StaticDir)
	} else {
		group.GET("/static/*filepath", func(c *gin.Context) {
			c.Request.URL.Path = constants.EmbedWebappPath + c.Param("filepath")
			http.FileServer(http.FS(constants.EmbedWebappFS)).ServeHTTP(c.Writer, c.Request)
		})
	}
}

func corsMiddleware() gin.HandlerFunc {
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
