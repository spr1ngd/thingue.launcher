package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thingue-launcher/common/constants"
	"thingue-launcher/common/provider"
)

type staticRouter struct{}

var StaticRouter = new(staticRouter)

func (s *staticRouter) BuildRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	Router.Static("/storage", constants.SAVE_DIR+"storage")
	if provider.AppConfig.LocalServer.UseExternalStatic {
		Router.Static("/customStatic", provider.AppConfig.LocalServer.StaticDir)
	}
	Router.GET("/static/*filepath", func(c *gin.Context) {
		c.Request.URL.Path = provider.WebStaticPath + c.Param("filepath")
		http.FileServer(http.FS(provider.WebStaticFiles)).ServeHTTP(c.Writer, c.Request)
	})
	return Router
}
