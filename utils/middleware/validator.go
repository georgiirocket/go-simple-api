package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-simple-api/utils/exception"
	"net/http"
)

func Validate[T any]() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data T

		errorReadJson := context.ShouldBindJSON(&data)

		if errorReadJson != nil {
			context.AbortWithStatusJSON(exception.NewBadRequestError("Read json"))

			return
		}

		validate := validator.New()

		validateErrors := validate.Struct(&data)

		if validateErrors != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": validateErrors.Error()})

			return
		}

		context.Set("data", data)
		context.Next()
	}
}
