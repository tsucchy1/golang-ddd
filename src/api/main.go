
package main

import (
	"github.com/gin-gonic/gin"
	"interfaces"
)

func main() {
		router := interfaces.AppRouter()

		router.GET("/ping", func(c *gin.Context){
			c.String("test")
		})

    router.Run(":8000")
}