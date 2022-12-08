package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-fiber-template/internal/infrastructure/database"
	"go-fiber-template/internal/infrastructure/mail"
	"go-fiber-template/internal/infrastructure/router"
	"go-fiber-template/internal/infrastructure/utils/vars"
	"log"
)

func Run() {
	mail.Connection(vars.MailConnection)

	// Connection database
	database.Connection()

	// Config API
	app := fiber.New(fiber.Config{
		AppName: vars.AppName,
	})
	app.Use(recover.New())
	router.Routes(app)

	// Run API
	log.Fatal(app.Listen(vars.PortAPI))
}
