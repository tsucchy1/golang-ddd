
package main

import (
	"github.com/gin-gonic/gin"
	"interface/router"
)

func main() {
		router := router.AppRouter()
    router.Run(":8000")
}