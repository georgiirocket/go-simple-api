package post

import (
	"github.com/gin-gonic/gin"
	"go-simple-api/utils/exception"
	"go-simple-api/utils/helpers"
	"go-simple-api/utils/schemas"
	"net/http"
)

type Controller struct {
	repository *Repository
}

func NewController(r *Repository) *Controller {
	return &Controller{
		repository: r,
	}
}

type CreatePostInput struct {
	Title       string `json:"title" validate:"required,min=5"`
	Description string `json:"description" validate:"required,min=5"`
}

// CreatePost
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param request body CreatePostInput true "data"
// @Success 201 {object} models.PostModel
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Security BearerAuth
// @Router       /api/post/create [post]
func (c *Controller) CreatePost(context *gin.Context) {
	data, err := helpers.GetContextData[CreatePostInput](context)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewBadRequestError(err.Error()))

		return
	}

	userId := helpers.GetContextUserId(context)

	post, err := c.repository.CreatePost(context, userId, data.Title, data.Description)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewServerError(err.Error()))

		return
	}

	context.JSON(http.StatusCreated, post.ToModel())
}

// GetPosts
// @Tags         Post
// @Accept       json
// @Produce      json
// @Success 200 {array} models.PostModel
// @Failure 500 {object} models.ErrorResponse
// @Security BearerAuth
// @Router       /api/post/list [get]
func (c *Controller) GetPosts(context *gin.Context) {
	posts, err := c.repository.GetPosts(context)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewServerError(err.Error()))

		return
	}

	context.JSON(http.StatusOK, schemas.PostsToModels(posts))
}

// GetPost
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        postId   path      string  true  "Post ID"
// @Success 200 {object} models.PostModel
// @Failure 404 {object} models.ErrorResponse
// @Security BearerAuth
// @Router       /api/post/view [get]
func (c *Controller) GetPost(context *gin.Context) {
	postId := context.Param("postId")
	post, err := c.repository.GetPostById(context, postId)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewNotFoundError(err.Error()))

		return
	}

	context.JSON(http.StatusOK, post.ToModel())
}

type UpdatePostInput struct {
	Title       string `json:"title" validate:"min=5"`
	Description string `json:"description" validate:"min=5"`
}

// UpdatePost
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        postId   path      string  true  "Post ID"
// @Param request body UpdatePostInput true "data"
// @Success 204 {object} models.PostModel
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Security BearerAuth
// @Router       /api/post/update [patch]
func (c *Controller) UpdatePost(context *gin.Context) {
	postId := context.Param("postId")

	post, err := c.repository.GetPostById(context, postId)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewNotFoundError(err.Error()))

		return
	}

	body, err := helpers.GetContextData[UpdatePostInput](context)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewBadRequestError(err.Error()))

		return
	}

	if len(body.Title) == 0 {
		body.Title = post.Title
	}

	if len(body.Description) == 0 {
		body.Description = post.Description
	}

	userId := helpers.GetContextUserId(context)

	updatedPost, err := c.repository.UpdatePost(context, userId, postId, body.Title, body.Description)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewServerError(err.Error()))

		return
	}

	context.JSON(http.StatusNoContent, updatedPost.ToModel())
}

// DeletePost
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        postId   path      string  true  "Post ID"
// @Success 200 {object} models.PostModel
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Security BearerAuth
// @Router       /api/post/delete [delete]
func (c *Controller) DeletePost(context *gin.Context) {
	postId := context.Param("postId")

	_, err := c.repository.GetPostById(context, postId)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewNotFoundError(err.Error()))

		return
	}

	userId := helpers.GetContextUserId(context)

	post, err := c.repository.DeletePost(context, userId, postId)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewServerError(err.Error()))

		return
	}

	context.JSON(http.StatusOK, post.ToModel())
}
