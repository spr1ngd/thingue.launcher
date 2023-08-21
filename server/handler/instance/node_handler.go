package instance

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"thingue-launcher/common/model"
	"thingue-launcher/common/model/response"
	"thingue-launcher/server/service"
)

func (g *HandlerGroup) NodeRegister(c *gin.Context) {
	var registerInfo model.NodeRegisterInfo
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
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)
}

func (g *HandlerGroup) UpdateProcessState(c *gin.Context) {
	var request model.ProcessStateUpdate
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	service.InstanceService.UpdateProcessState(&request)
	response.Ok(c)
}

func (g *HandlerGroup) ControlInstance(c *gin.Context) {

}
