package router

import (
	"github.com/gin-gonic/gin"
	"thingue-launcher/server/handler"
)

type instanceRouter struct{}

var InstanceRouter = new(instanceRouter)

func (s *instanceRouter) BuildRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	instanceRouter := Router.Group("instance")
	{
		instanceRouter.POST("nodeRegister", handler.InstanceGroup.NodeRegister)
		instanceRouter.GET("nodeList", handler.InstanceGroup.NodeList)
		instanceRouter.GET("getInstanceSid", handler.InstanceGroup.GetInstanceSid)
		instanceRouter.POST("instanceControl", handler.InstanceGroup.ControlInstance)
		instanceRouter.POST("getTicketByLabelSelector", handler.InstanceGroup.GetTicketByLabelSelector)
		instanceRouter.GET("getTicketById", handler.InstanceGroup.GetTicketById)
		instanceRouter.POST("updateProcessState", handler.InstanceGroup.UpdateProcessState)
	}
	return instanceRouter
}
