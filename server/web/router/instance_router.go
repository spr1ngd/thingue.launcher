package router

import (
	"github.com/gin-gonic/gin"
	"thingue-launcher/server/web/handler"
)

type instanceRouter struct{}

var InstanceRouter = new(instanceRouter)

func (s *instanceRouter) BuildRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	instanceRouter := Router.Group("instance")
	{
		instanceRouter.POST("clientRegister", handler.InstanceGroup.ClientRegister)
		instanceRouter.GET("clientList", handler.InstanceGroup.ClientList)
		instanceRouter.GET("getInstanceSid", handler.InstanceGroup.GetInstanceSid)
		instanceRouter.POST("processControl", handler.InstanceGroup.ProcessControl)
		instanceRouter.POST("pakControl", handler.InstanceGroup.PakControl)
		instanceRouter.POST("ticketSelect", handler.InstanceGroup.TicketSelect)
		instanceRouter.POST("instanceSelect", handler.InstanceGroup.InstanceSelect)
		instanceRouter.GET("getTicketById", handler.InstanceGroup.GetTicketById)
		instanceRouter.POST("updateProcessState", handler.InstanceGroup.UpdateProcessState)
		instanceRouter.POST("collectLogs", handler.InstanceGroup.CollectLogs)
		instanceRouter.POST("uploadLogs", handler.InstanceGroup.UploadLogs)
		instanceRouter.GET("downloadLogs", handler.InstanceGroup.DownloadLogs)
	}
	return instanceRouter
}
