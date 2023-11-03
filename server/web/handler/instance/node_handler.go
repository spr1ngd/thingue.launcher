package instance

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

func (g *HandlerGroup) NodeRegister(c *gin.Context) {
	var registerInfo request.NodeRegisterInfo
	err := c.ShouldBindJSON(&registerInfo)
	fmt.Printf("%+v\n", registerInfo)
	err = core.NodeService.NodeRegister(&registerInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("注册成功", c)
}

func (g *HandlerGroup) InstanceSelect(c *gin.Context) {
	var selectCond request.SelectorCond
	err := c.ShouldBindJSON(&selectCond)
	SelectedInstances, err := core.InstanceService.InstanceSelect(selectCond)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(SelectedInstances, c)
	}
}

func (g *HandlerGroup) GetInstanceSid(c *gin.Context) {
	nodeId := c.Query("nodeId")
	instanceId := c.Query("instanceId")
	sid, err := core.NodeService.GetInstanceSid(nodeId, instanceId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(sid, c)
}

func (g *HandlerGroup) NodeList(c *gin.Context) {
	list := core.NodeService.NodeList()
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "", c)
}

func (g *HandlerGroup) UpdateProcessState(c *gin.Context) {
	var msg message.NodeProcessStateUpdate
	err := c.ShouldBindJSON(&msg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	core.InstanceService.UpdateProcessState(&msg)
	response.Ok(c)
}

func (g *HandlerGroup) ProcessControl(c *gin.Context) {
	var req request.ProcessControl
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	core.InstanceService.ProcessControl(req)
	response.Ok(c)
}

func (g *HandlerGroup) PakControl(c *gin.Context) {
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

func (g *HandlerGroup) CollectLogs(c *gin.Context) {
	var req request.LogsCollect
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = core.NodeService.CollectLogs(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (g *HandlerGroup) UploadLogs(c *gin.Context) {
	traceId := c.Request.Header.Get("traceId")
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, c.Request.Body)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = core.NodeService.UploadLogs(traceId, buf)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (g *HandlerGroup) DownloadLogs(c *gin.Context) {
	traceId := c.Query("traceId")
	err, buf := core.NodeService.DownloadLogs(traceId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		c.Header("Content-Type", "application/zip")
		c.Header("Content-Disposition", "attachment; filename="+traceId+".zip")
		c.Writer.Write(buf.Bytes())
	}
}
