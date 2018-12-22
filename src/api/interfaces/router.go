
package interfaces

import (
	"github.com/gin-gonic/gin"
	"config/registry"
)

func AppRouter() {
	r := gin.Default()
	handler := registry.NewAppHandler()

	r.GET("/users/:name", handler.GetUser)
	
	return r
}