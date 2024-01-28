package rest

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
	"thingue-launcher/common/model"
	"thingue-launcher/common/request"
	"thingue-launcher/common/response"
	"thingue-launcher/common/util"
	"thingue-launcher/server/core"
	"thingue-launcher/server/core/provider"
)

type InstanceGroup struct{}

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

func (g *InstanceGroup) ClientList(c *gin.Context) {
	list := core.ClientService.ClientList()
	response.OkWithDetailed(response.PageResult[*model.Client]{
		List: list,
	}, "", c)
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
	err = core.InstanceService.PakControl(req.StreamerId, req.Type, req.PakName)
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
	ticket, err := core.TicketService.GetTicketById(c.Query("streamerId"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(ticket, c)
	}
}

func (g *InstanceGroup) InstanceList(c *gin.Context) {
	list := core.InstanceService.InstanceList()
	response.OkWithDetailed(response.PageResult[*model.Instance]{
		List: list,
	}, "", c)
}

func (g *InstanceGroup) GetInstanceByHostnameAndPid(c *gin.Context) {
	var instance *model.Instance
	hostname := c.Query("hostname")
	pidStr := c.Query("pid")
	pid, err := strconv.Atoi(pidStr)
	if err == nil {
		instance, err = core.InstanceService.GetInstanceByHostnameAndPid(hostname, pid)
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(instance, c)
	}
}

func (g *InstanceGroup) KickPlayerUser(c *gin.Context) {
	userQueryMap := map[string]string{}
	err := c.ShouldBindQuery(&userQueryMap)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	count, err := core.SdpService.KickPlayerUser(userQueryMap)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(count, fmt.Sprintf("踢掉%d个连接", count), c)
	}
}

func (g *InstanceGroup) SendMsgToStreamer(c *gin.Context) {
	var msg map[string]any
	err := c.BindJSON(&msg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		streamerId := c.Query("streamerId")
		streamer, err := provider.SdpConnProvider.GetStreamer(streamerId)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		}
		streamer.SendMessage(util.MapToJson(msg))
		response.OkWithMessage("消息发送成功", c)
	}
}
