package constants

type errorMessage struct {
	NotFoundErrorMessage   string
	BadRequestErrorMessage string
	ServerErrorMessage     string
	AuthErrorMessage       string
}

var ErrorMessages = errorMessage{
	NotFoundErrorMessage:   "Not found",
	BadRequestErrorMessage: "Bad request",
	ServerErrorMessage:     "Server error",
	AuthErrorMessage:       "Auth error",
}
