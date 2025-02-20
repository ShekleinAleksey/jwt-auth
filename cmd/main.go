package main

import (
	"log"
	"os"

	"github.com/ShekleinAleksey/jwt-auth/internal/handler"
	"github.com/ShekleinAleksey/jwt-auth/internal/repository"
	"github.com/ShekleinAleksey/jwt-auth/internal/service"
	"github.com/ShekleinAleksey/jwt-auth/pkg/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Auth Service
// @version 1.0
// @description API Service for Auth App
// @host http://95.174.91.82:8080
// @BasePath /
func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetOutput(os.Stdout)

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	logrus.Info("Initializing db...")

	db, err := postgres.NewDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	defer db.Close()
	if err != nil {
		log.Fatalf("error initializing db: %s", err.Error())
	}

	logrus.Info("Initializing repository...")
	repos := repository.NewRepository(db)
	logrus.Info("Initializing service...")
	services := service.NewService(repos)
	logrus.Info("Initializing handler...")
	handlers := handler.NewHandler(services)

	router := handlers.InitRoutes()

	logrus.Info("Starting server...")
	router.Run(":8080")
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
