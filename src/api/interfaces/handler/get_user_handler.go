
package handler

import (
	"github.com/gin-gonic/gin"
	"applicaton/usecase"
	"application/input"
	"config"
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
	input := &input.GetUser{name}
	output := h.u.GetUser(input)
	c.JSON(output)
}