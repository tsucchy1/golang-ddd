package handler

import (
	"api/application/input"
	"api/application/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetUserHandler interface {
	GetUser(c *gin.Context)
}

type getUserHandler struct {
	u usecase.UserUseCase
}

func NewGetUserHandler(u usecase.UserUseCase) GetUserHandler {
	return &getUserHandler{u}
}

func (h *getUserHandler) GetUser(c *gin.Context) {
	name := c.Param("name")
	input := input.GetUser{name}
	output, err := h.u.GetUser(input)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(200, output)
}
