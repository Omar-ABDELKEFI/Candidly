package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/database"
)

func Init() {
	app := fiber.New()
	database.GetDb()
	Router(app)
	database.MigrateDatabase()
	app.Listen(":8080")

}
