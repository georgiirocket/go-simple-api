package helpers

import (
	"github.com/gin-gonic/gin"
)

func GetContextUserId(context *gin.Context) string {
	data, isExist := context.Get("userId")

	if !isExist {
		return ""
	}

	return data.(string)
}
