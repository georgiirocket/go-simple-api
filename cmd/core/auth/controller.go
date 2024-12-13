package auth

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	repository *Repository
}

func NewController(r *Repository) *Controller {
	return &Controller{
		repository: r,
	}
}

type signInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Controller) SignUp(c *gin.Context) {
	//inp := new(signInput)
	//
	//if err := c.BindJSON(inp); err != nil {
	//	c.AbortWithStatus(http.StatusBadRequest)
	//	return
	//}
	//
	//if err := h.useCase.SignUp(c.Request.Context(), inp.Username, inp.Password); err != nil {
	//	c.AbortWithStatus(http.StatusInternalServerError)
	//	return
	//}
	//
	//c.Status(http.StatusOK)
}

type signInResponse struct {
	Token string `json:"token"`
}

func (h *Controller) SignIn(c *gin.Context) {
	//inp := new(signInput)
	//
	//if err := c.BindJSON(inp); err != nil {
	//	c.AbortWithStatus(http.StatusBadRequest)
	//	return
	//}
	//
	//token, err := h.useCase.SignIn(c.Request.Context(), inp.Username, inp.Password)
	//if err != nil {
	//	if err == auth.ErrUserNotFound {
	//		c.AbortWithStatus(http.StatusUnauthorized)
	//		return
	//	}
	//
	//	c.AbortWithStatus(http.StatusInternalServerError)
	//	return
	//}
	//
	//c.JSON(http.StatusOK, signInResponse{Token: token})
}
