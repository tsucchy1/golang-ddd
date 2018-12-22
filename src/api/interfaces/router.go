package interfaces

import (
	"api/interfaces/handler"
	"github.com/gin-gonic/gin"
)

func AppRouter(handler handler.AppHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/users/:name", handler.GetUser)

	return r
}
