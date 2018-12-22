package main

import (
	"api/config"
	"api/interfaces"
	"github.com/gin-gonic/gin"
)

func main() {
	registry := config.NewRegistry()
	router := interfaces.AppRouter(registry.NewAppHandler())

	router.Use(gin.Logger())
  router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})

	router.Run(":8000")
}
