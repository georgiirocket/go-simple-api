package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetContextData[T any](context *gin.Context) (*T, error) {
	data, isExist := context.Get("data")

	if !isExist {
		return nil, gin.Error{Err: http.ErrNoLocation, Type: gin.ErrorTypePrivate}
	}

	result, ok := data.(T)

	if !ok {
		return nil, gin.Error{Err: http.ErrBodyNotAllowed, Type: gin.ErrorTypeBind}
	}

	return &result, nil
}
