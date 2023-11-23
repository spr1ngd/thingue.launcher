package router

import (
	"github.com/gin-gonic/gin"
	"thingue-launcher/server/web/handler"
)

type mqttRouter struct{}

var MqttRouter = new(mqttRouter)

func (r *mqttRouter) BuildRouter(router *gin.RouterGroup) gin.IRoutes {
	return router.GET("/mqtt", func(c *gin.Context) {
		handler.MqttHandler.Handler(c.Writer, c.Request)
	})
}
