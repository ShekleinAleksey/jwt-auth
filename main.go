package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.GET("/request", handler.request)
	// r.GET("/refresh", handler.refresh)

	r.Run(":8080")
}
