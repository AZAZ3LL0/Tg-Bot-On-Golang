package main

import (
	"Telegram-Store/clients"
	"Telegram-Store/config"
	"Telegram-Store/handlers"
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

	log.Fatal(app.Listen(":" + config.Config("PORT")))
}
