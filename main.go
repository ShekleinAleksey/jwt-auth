package main

import (
	"log"

	"github.com/ShekleinAleksey/jwt-auth/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	r := gin.Default()
	r.POST("/create-token", handler.Token)
	r.POST("/refresh-token", handler.Refresh)

	r.Run(":8080")
}
