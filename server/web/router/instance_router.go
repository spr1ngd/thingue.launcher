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
		instanceRouter.POST("nodeRegister", handler.InstanceGroup.NodeRegister)
		instanceRouter.GET("nodeList", handler.InstanceGroup.NodeList)
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
