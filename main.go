package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ShekleinAleksey/jwt-auth/handler"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=localhost port=5432 user=admin dbname=jwt-auth password=%s sslmode=disable", os.Getenv("DB_PASSWORD")))
	if err != nil {
		log.Fatalf("failed to initialize db %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping %s", err.Error())
	}
	r := gin.Default()
	r.GET("/request", handler.Request)
	r.GET("/refresh", handler.Refresh)

	r.Run(":8080")
}
