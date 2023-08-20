package instance

import (
	"github.com/gin-gonic/gin"
	"thingue-launcher/common/model"
	"thingue-launcher/common/model/response"
	"thingue-launcher/server/service"
)

func (g *HandlerGroup) GetTicketByLabelSelector(c *gin.Context) {
	var selectCond model.SelectCond
	err := c.ShouldBindJSON(&selectCond)
	ticket, err := service.TicketService.GetTicketByLabelSelector(selectCond)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(ticket, c)
	}
}

func (g *HandlerGroup) GetTicketById(c *gin.Context) {
	ticket, err := service.TicketService.GetTicketById(c.Query("sid"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(ticket, c)
	}
}
