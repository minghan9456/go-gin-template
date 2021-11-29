package routes

import (
	"github.com/gin-gonic/gin"
	"oauth2-interface.pts.org.tw/src/controllers"
)

type Oauth2Route struct {
	Handler *gin.Engine
}

func NewOauth2Route(
	handler *gin.Engine,

) Oauth2Route {
	return Oauth2Route{
		Handler: handler,
	}
}

func (o Oauth2Route) Setup() {
	oauth2Ctrl := controllers.NewOauth2Controller()

	post := o.Handler.Group("/oauth2")
	{
		post.GET("/auth", oauth2Ctrl.Authorize)
		post.POST("/token", oauth2Ctrl.GenToken)
		post.POST("/revoke", oauth2Ctrl.RevokeToken)
	}
}
