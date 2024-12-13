package exception

import (
	"github.com/gin-gonic/gin"
	"go-simple/src/constants"
	"go-simple/src/dto"
	"net/http"
)

func NotFound(context *gin.Context, message string) {
	errorResponse := dto.ErrorResponse{
		Error: constants.NotFoundErrorMessage,
	}

	if len(message) == 0 {
		context.JSON(http.StatusFound, errorResponse)

		return
	}

	errorResponse.Error += ": " + message

	context.JSON(http.StatusFound, errorResponse)
}

func ServerError(context *gin.Context, message string) {
	errorResponse := dto.ErrorResponse{
		Error: constants.ServerErrorMessage,
	}

	if len(message) == 0 {
		context.JSON(http.StatusInternalServerError, errorResponse)

		return
	}

	errorResponse.Error += ": " + message

	context.JSON(http.StatusInternalServerError, errorResponse)
}

func BadRequestError(context *gin.Context, message string) {
	errorResponse := dto.ErrorResponse{
		Error: constants.BadRequestErrorMessage,
	}

	if len(message) == 0 {
		context.JSON(http.StatusBadRequest, errorResponse)

		return
	}

	errorResponse.Error += ": " + message

	context.JSON(http.StatusBadRequest, errorResponse)
}

func AuthError(context *gin.Context, message string) {
	errorResponse := dto.ErrorResponse{
		Error: constants.AuthErrorMessage,
	}

	if len(message) == 0 {
		context.JSON(http.StatusUnauthorized, errorResponse)

		return
	}

	errorResponse.Error += ": " + message

	context.JSON(http.StatusUnauthorized, errorResponse)
}
