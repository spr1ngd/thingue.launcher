package rest

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"thingue-launcher/common/message"
	"thingue-launcher/common/request"
	"thingue-launcher/common/response"
	"thingue-launcher/server/core"
)

type InstanceGroup struct{}

func (g *InstanceGroup) ClientRegister(c *gin.Context) {
	var registerInfo request.ClientRegisterInfo
	err := c.ShouldBindJSON(&registerInfo)
	fmt.Printf("%+v\n", registerInfo)
	err = core.ClientService.ClientRegister(&registerInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("注册成功", c)
}

func (g *InstanceGroup) InstanceSelect(c *gin.Context) {
	var selectCond request.SelectorCond
	err := c.ShouldBindJSON(&selectCond)
	SelectedInstances, err := core.InstanceService.InstanceSelect(selectCond)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(SelectedInstances, c)
	}
}

func (g *InstanceGroup) GetInstanceSid(c *gin.Context) {
	clientId := c.Query("clientId")
	instanceId := c.Query("instanceId")
	sid, err := core.ClientService.GetInstanceSid(clientId, instanceId)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
		return
	}
	response.OkWithData(sid, c)
}

func (g *InstanceGroup) ClientList(c *gin.Context) {
	list := core.ClientService.ClientList()
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "", c)
}

func (g *InstanceGroup) UpdateProcessState(c *gin.Context) {
	var msg message.ClientProcessStateUpdate
	err := c.ShouldBindJSON(&msg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	core.InstanceService.UpdateProcessState(&msg)
	response.Ok(c)
}

func (g *InstanceGroup) ProcessControl(c *gin.Context) {
	var req request.ProcessControl
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	core.InstanceService.ProcessControl(req)
	response.Ok(c)
}

func (g *InstanceGroup) PakControl(c *gin.Context) {
	var req request.PakControl
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = core.InstanceService.PakControl(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (g *InstanceGroup) CollectLogs(c *gin.Context) {
	var req request.LogsCollect
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = core.ClientService.CollectLogs(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (g *InstanceGroup) UploadLogs(c *gin.Context) {
	traceId := c.Request.Header.Get("traceId")
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, c.Request.Body)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = core.ClientService.UploadLogs(traceId, buf)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (g *InstanceGroup) DownloadLogs(c *gin.Context) {
	traceId := c.Query("traceId")
	err, buf := core.ClientService.DownloadLogs(traceId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		c.Header("Content-Type", "application/zip")
		c.Header("Content-Disposition", "attachment; filename="+traceId+".zip")
		c.Writer.Write(buf.Bytes())
	}
}

func (g *InstanceGroup) TicketSelect(c *gin.Context) {
	var selectCond request.SelectorCond
	err := c.ShouldBindJSON(&selectCond)
	ticket, err := core.TicketService.TicketSelect(selectCond)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(ticket, c)
	}
}

func (g *InstanceGroup) GetTicketById(c *gin.Context) {
	ticket, err := core.TicketService.GetTicketById(c.Query("sid"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(ticket, c)
	}
}

func (g *InstanceGroup) InstanceList(c *gin.Context) {
	list, _ := core.InstanceService.InstanceList()
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "", c)
}
