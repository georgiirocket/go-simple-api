package auth

import (
	"github.com/gin-gonic/gin"
	"go-simple-api/utils/middleware"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, repository *Repository) {
	controller := NewController(repository)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/sign-up", middleware.Validate[SignInput](), controller.SignUp)
		authEndpoints.POST("/sign-in", middleware.Validate[SignInput](), controller.SignIn)
		authEndpoints.POST("/refresh-tokens", middleware.Validate[RefreshInput](), controller.RefreshToken)
	}
}
