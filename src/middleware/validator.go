package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-simple/src/exception"
	"net/http"
)

func ValidateBy[T any]() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data T

		errorReadJson := context.ShouldBindJSON(&data)

		if errorReadJson != nil {
			exception.BadRequestError(context, "Read json")
			context.Abort()

			return
		}

		validate := validator.New()

		validateErrors := validate.Struct(&data)

		if validateErrors != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": validateErrors.Error()})
			context.Abort()

			return
		}

		context.Set("data", data)
		context.Next()
	}
}
