package middleware

import (
	"github.com/gin-gonic/gin"
	"go-simple-api/utils/exception"
	"go-simple-api/utils/services"
	"strings"
)

func AccessGuard(context *gin.Context) {
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

	authData, err := services.VerifyToken(token[1], services.AccessTokenType)

	if err != nil {
		context.AbortWithStatusJSON(exception.NewUnauthorizedError(err.Error()))

		return
	}

	context.Set("userId", authData.UserId)
	context.Next()
}
