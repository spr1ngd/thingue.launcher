package agent

import (
	"github.com/gin-gonic/gin"
	"thingue-launcher/common/model"
	"thingue-launcher/common/model/response"
)

type BaseApi struct{}

func (b *BaseApi) Register(c *gin.Context) {
	var registerInfo model.AgentRegisterInfo
	err := c.ShouldBindJSON(&registerInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
}
