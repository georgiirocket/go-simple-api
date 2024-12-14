package auth

import (
	"github.com/gin-gonic/gin"
	"go-simple-api/utils/exception"
	"go-simple-api/utils/helpers"
	"go-simple-api/utils/models"
	"go-simple-api/utils/services"
	"net/http"
	"strings"
)

type Controller struct {
	repository *Repository
}

func NewController(r *Repository) *Controller {
	return &Controller{
		repository: r,
	}
}

type SignInput struct {
	Username string `json:"username" validate:"required,min=5"`
	Password string `json:"password" validate:"required,min=10"`
}

// SignUp
// @Summary get new  tokens
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param request body SignInput true "data"
// @Success 200 {object} models.UserModel
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Security BearerAuth
// @Router       /api/auth/sign-up [post]
func (c *Controller) SignUp(context *gin.Context) {
	data, err := helpers.GetContextData[SignInput](context)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewBadRequestError(err.Error()))

		return
	}

	user, _ := c.repository.GetUser(context, data.Username)

	if user != nil {
		context.AbortWithStatusJSON(exception.NewBadRequestError("Chose another username"))

		return
	}

	hash, err := services.CreateHashPassword(data.Password)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewServerError("Create password"))
	}

	newUser, err := c.repository.CreateUser(context, data.Username, hash)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewServerError("When creating user"))

		return
	}

	context.JSON(http.StatusCreated, newUser.ToModel())
}

type SignInResponse struct {
	User models.UserModel `json:"user"`
	Auth models.AuthData  `json:"auth"`
}

func (c *Controller) SignIn(context *gin.Context) {
	data, err := helpers.GetContextData[SignInput](context)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewBadRequestError("Read json"))

		return
	}

	user, _ := c.repository.GetUser(context, data.Username)

	if user == nil {
		context.AbortWithStatusJSON(exception.NewUnauthorizedError("Check your credentials"))

		return
	}

	isVerifyPassword := services.VerifyPassword(data.Password, user.Password)

	if !isVerifyPassword {
		context.AbortWithStatusJSON(exception.NewUnauthorizedError("Check your credentials"))

		return
	}

	payload, err := services.CreateAuthData(user.ID.Hex())

	if err != nil {
		context.AbortWithStatusJSON(exception.NewServerError(err.Error()))

		return
	}

	context.JSON(http.StatusCreated, SignInResponse{Auth: *payload, User: user.ToModel()})
}

type RefreshInput struct {
	AccessToken string `json:"accessToken" validate:"required,min=10"`
}

func (c *Controller) RefreshToken(context *gin.Context) {
	data, err := helpers.GetContextData[RefreshInput](context)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewBadRequestError("Read json"))

		return
	}

	authorization := context.GetHeader("Authorization")

	if authorization == "" {
		context.AbortWithStatusJSON(exception.NewUnauthorizedError("Unauthorized"))

		return
	}

	token := strings.Split(authorization, " ")

	if len(token) != 2 || token[0] != "Bearer" {
		context.AbortWithStatusJSON(exception.NewUnauthorizedError("Invalid authorization header format"))

		return
	}

	payload, err := services.VerifyRefreshToken(token[1], data.AccessToken)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewUnauthorizedError(err.Error()))

		return
	}

	user, err := c.repository.GetUserById(context, payload.UserId)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewUnauthorizedError("Unauthorized"))

		return
	}

	newPayload, err := services.CreateAuthData(user.ID.Hex())

	if err != nil {
		context.AbortWithStatusJSON(exception.NewServerError(err.Error()))

		return
	}

	context.JSON(http.StatusCreated, newPayload)
}
