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

	// db, err := postgres.NewDB(postgres.Config{
	// 	Host:     viper.GetString("db.host"),
	// 	Port:     viper.GetString("db.port"),
	// 	Username: viper.GetString("db.username"),
	// 	DBName:   viper.GetString("db.dbname"),
	// 	SSLMode:  viper.GetString("db.sslmode"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// })

	r := gin.Default()
	r.POST("/create-user", handler.CreateUser)
	r.GET("/get-user", handler.GetUser)
	r.POST("/create-token", handler.Token)
	r.POST("/refresh-token", handler.Refresh)

	r.Run(":8080")
}
