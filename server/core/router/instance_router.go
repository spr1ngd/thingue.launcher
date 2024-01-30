package router

import (
	"github.com/gin-gonic/gin"
	"thingue-launcher/server/core/handler"
)

type instanceRouter struct{}

var InstanceRouter = new(instanceRouter)

func (s *instanceRouter) BuildRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	instanceRouter := Router.Group("instance")
	{
		instanceRouter.GET("clientList", handler.InstanceGroup.ClientList)
		instanceRouter.POST("processControl", handler.InstanceGroup.ProcessControl)
		instanceRouter.POST("pakControl", handler.InstanceGroup.PakControl)
		instanceRouter.POST("sendMsgToStreamer", handler.InstanceGroup.SendMsgToStreamer)
		instanceRouter.POST("ticketSelect", handler.InstanceGroup.TicketSelect)
		instanceRouter.GET("instanceList", handler.InstanceGroup.InstanceList)
		instanceRouter.POST("instanceSelect", handler.InstanceGroup.InstanceSelect)
		instanceRouter.GET("getTicketById", handler.InstanceGroup.GetTicketById)
		instanceRouter.GET("collectLogs", handler.InstanceGroup.CollectLogs)
		instanceRouter.GET("getInstanceByHostnameAndPid", handler.InstanceGroup.GetInstanceByHostnameAndPid)
		instanceRouter.GET("kickPlayerUser", handler.InstanceGroup.KickPlayerUser)
	}
	return instanceRouter
}
