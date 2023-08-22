package instance

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"thingue-launcher/common/message"
	"thingue-launcher/common/request"
	"thingue-launcher/common/response"
	"thingue-launcher/server/service"
)

func (g *HandlerGroup) NodeRegister(c *gin.Context) {
	var registerInfo request.NodeRegisterInfo
	err := c.ShouldBindJSON(&registerInfo)
	fmt.Printf("%+v\n", registerInfo)
	err = service.NodeService.NodeRegister(&registerInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("注册成功", c)
}

func (g *HandlerGroup) GetInstanceSid(c *gin.Context) {
	nodeId := c.Query("nodeId")
	instanceId := c.Query("instanceId")
	sid, err := service.NodeService.GetInstanceSid(nodeId, instanceId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(sid, c)
}

func (g *HandlerGroup) NodeList(c *gin.Context) {
	list := service.NodeService.NodeList()
	response.OkWithData(response.PageResult{
		List: list,
	}, c)
}

func (g *HandlerGroup) UpdateProcessState(c *gin.Context) {
	var request message.ProcessStateUpdate
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	service.InstanceService.UpdateProcessState(&request)
	response.Ok(c)
}

func (g *HandlerGroup) ProcessControl(c *gin.Context) {
	var req request.ProcessControl
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	service.InstanceService.ProcessControl(req)
	response.Ok(c)
}
