package post

import (
	"github.com/gin-gonic/gin"
	"go-simple-api/utils/middleware"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, repository *Repository) {
	controller := NewController(repository)

	authEndpoints := router.Group("/post")
	authEndpoints.Use(middleware.AccessGuard)

	{
		authEndpoints.POST("/create", middleware.Validate[CreatePostInput](), controller.CreatePost)
		authEndpoints.GET("/list", controller.GetPosts)
		authEndpoints.GET("/view/:postId", controller.GetPost)
		authEndpoints.PATCH("/update/:postId", middleware.Validate[UpdatePostInput](), controller.UpdatePost)
		authEndpoints.DELETE("/delete/:postId", controller.DeletePost)
	}
}
