package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//GinRouter -> Gin Router
type GinRouter struct {
	Gin *gin.Engine
}

//NewGinRouter all the routes are defined here
func NewGinRouter() GinRouter {

	httpRouter := gin.Default()

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})

	postRoute := NewOauth2Route(httpRouter) // oauth2 routes are initialized
	postRoute.Setup()                       // oauth2 routes are being setup

	return GinRouter{
		Gin: httpRouter,
	}

}
