package router

import (
	"github.com/gin-gonic/gin"
	"thingue-launcher/server/web/handler"
)

type wsRouter struct{}

var WsRouter = new(wsRouter)

func (s *wsRouter) BuildRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	wsRouter := Router.Group("/ws")
	{
		wsRouter.GET("/streamer/:id", handler.WsGroup.StreamerWebSocketHandler)
		wsRouter.GET("/player/:ticket", handler.WsGroup.PlayerWebSocketHandler)
		wsRouter.GET("/client", handler.WsGroup.ClientWebSocketHandler)
		wsRouter.GET("/admin", handler.WsGroup.AdminWebSocketHandler)
	}
	return wsRouter
}
