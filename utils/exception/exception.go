package exception

import (
	"go-simple-api/utils/constants"
	"go-simple-api/utils/models"
	"net/http"
)

func NewServerError(message string) (int, models.ErrorResponse) {
	res := models.ErrorResponse{Error: constants.ErrorMessages.ServerErrorMessage}

	if len(message) > 0 {
		res.Error += ": " + message
	}

	return http.StatusInternalServerError, res
}

func NewNotFoundError(message string) (int, models.ErrorResponse) {
	res := models.ErrorResponse{Error: constants.ErrorMessages.NotFoundErrorMessage}

	if len(message) > 0 {
		res.Error += ": " + message
	}

	return http.StatusNotFound, res
}

func NewBadRequestError(message string) (int, models.ErrorResponse) {
	res := models.ErrorResponse{Error: constants.ErrorMessages.BadRequestErrorMessage}

	if len(message) > 0 {
		res.Error += ": " + message
	}

	return http.StatusBadRequest, res
}

func NewUnauthorizedError(message string) (int, models.ErrorResponse) {
	res := models.ErrorResponse{Error: constants.ErrorMessages.AuthErrorMessage}

	if len(message) > 0 {
		res.Error += ": " + message
	}

	return http.StatusUnauthorized, res
}
