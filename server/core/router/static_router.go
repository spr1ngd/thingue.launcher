package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/provider"
)

type staticRouter struct{}

var StaticRouter = new(staticRouter)

func (s *staticRouter) BuildRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	Router.Static("/storage", path.Join(constants.SaveDir, "storage"))
	if provider.AppConfig.LocalServer.UseExternalStatic {
		Router.Static("/customStatic", provider.AppConfig.LocalServer.StaticDir)
	}
	Router.GET("/static/*filepath", func(c *gin.Context) {
		c.Request.URL.Path = constants.EmbedWebappPath + c.Param("filepath")
		http.FileServer(http.FS(constants.EmbedWebappFS)).ServeHTTP(c.Writer, c.Request)
	})
	return Router
}
