package app

import (
	"log"
	"net/http"
	"os"

	"github.com/ShekleinAleksey/goTickets/internal/handler"
	"github.com/ShekleinAleksey/goTickets/internal/repository"
	"github.com/ShekleinAleksey/goTickets/internal/service"
	"github.com/ShekleinAleksey/goTickets/pkg/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Tickets Service
// @version 1.0
// @description API Service for Tickets App

// @host localhost:8080
// @BasePath /api/v1/
func Run() {
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
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	logrus.Info("Initializing repository...")
	repos := repository.NewRepository(db)
	logrus.Info("Initializing service...")
	services := service.NewService(repos)
	logrus.Info("Initializing handler...")
	handlers := handler.NewHandler(services)

	router := handlers.InitRoutes()

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", router)
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
