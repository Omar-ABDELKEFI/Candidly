package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tekab-dev/tekab-test/database"
)

func Init() {
	app := fiber.New()
	app.Use(cors.New())
	database.GetDb()
	Router(app)
	database.MigrateDatabase()
	app.Listen(":8080")

}
