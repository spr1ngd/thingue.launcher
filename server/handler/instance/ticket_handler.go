package instance

import (
	"github.com/gin-gonic/gin"
	"thingue-launcher/common/request"
	"thingue-launcher/common/response"
	"thingue-launcher/server/service"
)

func (g *HandlerGroup) TicketSelect(c *gin.Context) {
	var selectCond request.TicketSelector
	err := c.ShouldBindJSON(&selectCond)
	ticket, err := service.TicketService.TicketSelect(selectCond)
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
