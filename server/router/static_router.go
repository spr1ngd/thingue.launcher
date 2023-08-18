package router

import (
	"embed"
	"github.com/gin-gonic/gin"
	"net/http"
)

type staticRouter struct{}

var StaticRouter = new(staticRouter)

func (s *staticRouter) BuildRouter(Router *gin.RouterGroup, staticFiles embed.FS) (R gin.IRoutes) {
	Router.GET("/static/*filepath", func(c *gin.Context) {
		c.Request.URL.Path = "/frontend/dist" + c.Param("filepath")
		http.FileServer(http.FS(staticFiles)).ServeHTTP(c.Writer, c.Request)
	})
	return Router
}
