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

func (g *HandlerGroup) NodeList(c *gin.Context) {
	list := service.NodeService.NodeList()
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)
}

func (g *HandlerGroup) UpdateInstance(c *gin.Context) {

}

func (g *HandlerGroup) ControlInstance(c *gin.Context) {

}
