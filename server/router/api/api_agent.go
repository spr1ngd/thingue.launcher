package api

import (
	"github.com/gin-gonic/gin"
	v1 "thingue-launcher/server/api/v1"
)

type AgentRouter struct{}

func (s *AgentRouter) InitAgentRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	agentRouter := Router.Group("agent")
	baseApi := v1.ApiGroupApp.Agent.BaseApi
	{
		agentRouter.POST("register", baseApi.Register)
	}
	return agentRouter
}
