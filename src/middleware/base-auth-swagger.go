package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
)

func BaseAuthSwagger() gin.HandlerFunc {
	user := os.Getenv("SWAGGER_USER")
	password := os.Getenv("SWAGGER_PASSWORD")

	return gin.BasicAuth(gin.Accounts{user: password})
}
