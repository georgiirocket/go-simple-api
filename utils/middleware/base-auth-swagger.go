package middleware

import (
	"github.com/gin-gonic/gin"
	"go-simple-api/config"
)

func BaseAuthSwagger() gin.HandlerFunc {
	user := config.Env.SwaggerUser
	password := config.Env.SwaggerPassword

	return gin.BasicAuth(gin.Accounts{user: password})
}
