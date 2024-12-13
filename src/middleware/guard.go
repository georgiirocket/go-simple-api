package middleware

import (
	"github.com/gin-gonic/gin"
	"go-simple/src/exception"
	"go-simple/src/services"
	"strings"
)

func AccessGuard(context *gin.Context) {
	authorization := context.GetHeader("Authorization")

	if authorization == "" {
		exception.AuthError(context, "Unauthorized")
		context.Abort()

		return
	}

	token := strings.Split(authorization, " ")

	if len(token) != 2 || token[0] != "Bearer" {
		exception.AuthError(context, "Invalid authorization header format")
		context.Abort()

		return
	}

	authData, err := services.VerifyToken(token[1], services.AccessTokenType)

	if err != nil {
		exception.AuthError(context, err.Error())
		context.Abort()

		return
	}

	context.Set("userId", authData.UserId)
	context.Next()
}
