
package interfaces

import (
	"github.com/gin-gonic/gin"
	"api/interfaces/handler"
)

func AppRouter(handler handler.AppHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/users/:name", handler.GetUser)
	
	return r
}