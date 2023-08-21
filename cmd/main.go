package main

import (
	"Telegram-Store/internal/clients"
	"Telegram-Store/internal/handlers"
	"Telegram-Store/internal/storage/database"
	"Telegram-Store/internal/storage/repositories"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func initConfig() error {
	viper.AddConfigPath("cmd/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	fmt.Println("Проверка 1")
	if err := initConfig(); err != nil {
		log.Fatalf("error occured load config file: %s", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error occured load .env file: %s", err.Error())
	}
	fmt.Println("Начинается подключение к бд")
	// Подключение к базе данных
	db, err := database.NewPostgresDB(&database.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.db_name"),
		SslMode:  viper.GetString("db.ssl_mode"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Успешно подключились к бд")

	defer db.Close()

	// Создание репозитория
	bookRepo := repositories.NewbookRepository(db)
	userRepo := repositories.NewUserRepository(db)
	fmt.Println("Создание репозитория")

	// Создание бота
	bot := clients.Init()
	fmt.Println("Создание бота")

	// Инициализация обработчиков
	ctx := context.Background()
	fmt.Println("Инициализация обработчиков1")

	handlers.Init(ctx, bot, bookRepo, userRepo)

}
