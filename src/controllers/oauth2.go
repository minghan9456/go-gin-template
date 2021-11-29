package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Oauth2Controller struct {
	// service service.OauthService
}

func NewOauth2Controller() Oauth2Controller {
	return Oauth2Controller{}
}

/*
func NewOauth2Controller(s service.OauthService) Oauth2Controller {
	return Oauth2Controller{
		service: s,
	}
}
*/

func (p Oauth2Controller) Authorize(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Authorize API"})
}

func (p Oauth2Controller) GenToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "GenToken API"})
}

func (p Oauth2Controller) RevokeToken(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "RevokeToken API"})
}
