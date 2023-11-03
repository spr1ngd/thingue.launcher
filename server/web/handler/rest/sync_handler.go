package rest

import (
	"github.com/gin-gonic/gin"
	"thingue-launcher/common/model"
	"thingue-launcher/common/response"
	"thingue-launcher/server/core"
)

type SyncGroup struct{}

func (g *SyncGroup) GetSyncConfig(c *gin.Context) {
	config := core.SyncService.GetSyncConfig()
	response.OkWithData(config, c)
}

func (g *SyncGroup) GetCloudFiles(c *gin.Context) {
	resName := c.Query("resName")
	files := core.SyncService.GetCloudFiles(resName)
	response.OkWithData(files, c)
}

func (g *SyncGroup) UpdateCloudFiles(c *gin.Context) {
	resName := c.Query("resName")
	var files []*model.CloudFile
	_ = c.ShouldBindJSON(&files)
	core.SyncService.UpdateCloudFiles(resName, files)
	response.Ok(c)
}

func (g *SyncGroup) UploadFile(c *gin.Context) {
	path := c.Request.Header.Get("path")
	res := c.Request.Header.Get("res")
	err := core.SyncService.UploadFile(res, path, c.Request.Body)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
