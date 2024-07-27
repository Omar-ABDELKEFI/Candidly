package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/dev/test/database"
)

func Init() {
	app := fiber.New()
	// enable Cross-Origin
	app.Use(cors.New())
	database.GetDb()
	Router(app)
	app.Listen(":8080")

}
