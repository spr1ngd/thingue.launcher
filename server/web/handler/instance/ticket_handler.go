package instance

import (
	"github.com/gin-gonic/gin"
	"thingue-launcher/common/request"
	"thingue-launcher/common/response"
	"thingue-launcher/server/core"
)

func (g *HandlerGroup) TicketSelect(c *gin.Context) {
	var selectCond request.SelectorCond
	err := c.ShouldBindJSON(&selectCond)
	ticket, err := core.TicketService.TicketSelect(selectCond)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(ticket, c)
	}
}

func (g *HandlerGroup) GetTicketById(c *gin.Context) {
	ticket, err := core.TicketService.GetTicketById(c.Query("sid"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(ticket, c)
	}
}
