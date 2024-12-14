package auth

import (
	"github.com/gin-gonic/gin"
	"go-simple-api/utils/exception"
	"go-simple-api/utils/helpers"
	"go-simple-api/utils/models"
	"go-simple-api/utils/services"
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

type SignInput struct {
	Username string `json:"username" validate:"required,min=5"`
	Password string `json:"password" validate:"required,min=10"`
}

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

	authData, errAuth := services.CreateAuthData(user.ID.Hex())

	if errAuth != nil {
		context.AbortWithStatusJSON(exception.NewServerError(errAuth.Error()))

		return
	}

	context.JSON(http.StatusCreated, SignInResponse{Auth: *authData, User: user.ToModel()})
}
