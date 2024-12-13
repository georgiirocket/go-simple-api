package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, repository *Repository) {
	controller := NewController(repository)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/sign-up", controller.SignUp)
		authEndpoints.POST("/sign-in", controller.SignIn)
	}
}
