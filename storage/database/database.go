package database

import (
	"Telegram-Store/cmd/config"
	"Telegram-Store/storage/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s",
		config.Config("POSTGRES_HOST"),
		config.Config("POSTGRES_USER"),
		config.Config("POSTGRES_PASSWORD"),
		config.Config("POSTGRES_PORT"),
		config.Config("POSTGRES_DATABASE_NAME"),
		config.Config("SSL_MODE"),
	)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if err = DB.AutoMigrate(&models.Book{}); err != nil {
		fmt.Println(err)
		return nil
	}

	return DB
}
