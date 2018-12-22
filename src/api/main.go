
package main

import (
	"github.com/gin-gonic/gin"
	"api/interfaces"
	"api/config"
)

func main() {
		registry := config.NewRegistry()
		router := interfaces.AppRouter(registry.NewAppHandler())

		router.GET("/ping", func(c *gin.Context){
			c.String(200, "test")
		})

    router.Run(":8000")
}