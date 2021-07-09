package server

import (
	"github.com/gofiber/fiber/v2"
)

func Init() {
	app := fiber.New()
	Router(app)
	app.Listen(":8080")

}
