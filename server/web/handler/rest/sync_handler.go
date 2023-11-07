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
	res := c.Query("res")
	var files []*model.CloudFile
	_ = c.ShouldBindJSON(&files)
	core.SyncService.UpdateCloudFiles(res, files)
	response.Ok(c)
}

func (g *SyncGroup) DeleteCloudFiles(c *gin.Context) {
	res := c.Query("res")
	var names []string
	_ = c.ShouldBindJSON(&names)
	core.SyncService.DeleteFiles(res, names)
	response.Ok(c)
}
func (g *SyncGroup) UploadFile(c *gin.Context) {
	name := c.Request.Header.Get("name")
	res := c.Request.Header.Get("res")
	err := core.SyncService.UploadFile(res, name, c.Request.Body)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
