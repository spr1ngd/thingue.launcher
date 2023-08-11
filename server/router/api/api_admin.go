package api

import (
	"github.com/gin-gonic/gin"
	v1 "thingue-launcher/server/api/v1"
)

type AdminRouter struct{}

func (s *AgentRouter) InitAdminRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	adminRouter := Router.Group("admin")
	baseApi := v1.ApiGroupApp.Admin.BaseApi
	{
		adminRouter.POST("register", baseApi.Register)
	}
	return adminRouter
}
