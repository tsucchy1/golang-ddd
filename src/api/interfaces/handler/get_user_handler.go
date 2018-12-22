
package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"api/application/usecase"
	"api/application/input"
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

	if err := c.ShouldBindJSON(output); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	c.JSON(200, gin.H{"status": "Success"})
}