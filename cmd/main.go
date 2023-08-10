package main

import (
	"Telegram-Store/cmd/config"
	"Telegram-Store/internal/clients"
	"Telegram-Store/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	bot := clients.Init()
	handlers.Init(bot)

	port, err := config.Config("PORT")

	if err != nil {
		log.Println(err)
		return
	}

	log.Print(app.Listen(":" + port))
}
