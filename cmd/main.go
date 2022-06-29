package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"restaurant-reservation/internal/handler/rest"
	"restaurant-reservation/internal/repository"
	"restaurant-reservation/internal/repository/postgres"
	"restaurant-reservation/internal/service"
)

func main() {
	// инициализирую yml конфиг
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	// инициализирую env конфиг
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	//соединение с бд
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db %s", err.Error())
	}

	//прокидываю инстанс бд и создаю репозитории
	repos := repository.NewRepository(db)
	//прокидываю репозитории в сервисы
	services := service.NewService(repos)
	//сервисы в хендлеры
	handlers := rest.NewHandler(services)

	//запускаю сервер на порту 8000
	srv := new(Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while runnning http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigFile("config")
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}
