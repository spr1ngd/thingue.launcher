package router

import (
	"github.com/gin-gonic/gin"
	"thingue-launcher/server/core/handler"
)

type syncRouter struct{}

var SyncRouter = new(syncRouter)

func (s *syncRouter) BuildRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	syncRouter := Router.Group("sync")
	{
		syncRouter.GET("getSyncConfig", handler.SyncGroup.GetSyncConfig)
		syncRouter.GET("listCloudRes", handler.SyncGroup.ListCloudRes)
		syncRouter.POST("updateCloudRes", handler.SyncGroup.UpdateCloudRes)
		syncRouter.POST("createCloudRes", handler.SyncGroup.CreateCloudRes)
		syncRouter.POST("deleteCloudRes", handler.SyncGroup.DeleteCloudRes)
		syncRouter.GET("getCloudFiles", handler.SyncGroup.GetCloudFiles)
		syncRouter.POST("updateCloudFiles", handler.SyncGroup.UpdateCloudFiles)
		syncRouter.POST("deleteCloudFiles", handler.SyncGroup.DeleteCloudFiles)
		syncRouter.POST("uploadFile", handler.SyncGroup.UploadFile)
	}
	return syncRouter
}
