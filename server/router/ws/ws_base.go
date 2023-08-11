package ws

import (
	"github.com/gin-gonic/gin"
	"thingue-launcher/server/ws"
)

type BaseRouter struct{}

func (s *BaseRouter) InitWsRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	wsRouter := Router.Group("/ws")
	{
		wsRouter.GET("/streamer/:id", ws.StreamerWebSocketHandler)
		wsRouter.GET("/player/:streamerId", ws.PlayerWebSocketHandler)
		wsRouter.GET("/agent", ws.AgentWebSocketHandler)
		wsRouter.GET("/admin", ws.AdminWebSocketHandler)
	}
	return wsRouter
}
