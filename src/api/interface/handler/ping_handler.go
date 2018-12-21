
package handler

import "github.com/gin-gonic/gin"

type PingHandler interface {
}

type pingHandler struct {
}

func (h *pingHandler) Ping(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
}