package router

import (
	"github.com/gin-gonic/gin"
	"thingue-launcher/server/web/handler"
)

type syncRouter struct{}

var SyncRouter = new(syncRouter)

func (s *syncRouter) BuildRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	syncRouter := Router.Group("sync")
	{
		syncRouter.GET("getSyncConfig", handler.SyncGroup.GetSyncConfig)
		syncRouter.GET("getCloudFiles", handler.SyncGroup.GetCloudFiles)
		syncRouter.POST("updateCloudFiles", handler.SyncGroup.UpdateCloudFiles)
		syncRouter.POST("deleteCloudFiles", handler.SyncGroup.DeleteCloudFiles)
		syncRouter.POST("uploadFile", handler.SyncGroup.UploadFile)
	}
	return syncRouter
}
