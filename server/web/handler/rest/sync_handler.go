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
	res := c.Query("res")
	files := core.SyncService.GetCloudFiles(res)
	response.OkWithData(files, c)
}

func (g *SyncGroup) UpdateCloudFiles(c *gin.Context) {
	//res := c.Query("res")
	var files []*model.CloudFile
	_ = c.ShouldBindJSON(&files)
	core.SyncService.UpdateCloudFiles(files)
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
