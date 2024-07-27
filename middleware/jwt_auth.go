package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/dev/test/services"
	"strings"
)

func TokenAuthMiddleware() func(ctx *fiber.Ctx) error {

	return func(ctx *fiber.Ctx) error {

		authHeader := ctx.Get("Authorization")

		if authHeader == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "token invalid",
			})
		}

		splitToken := strings.Split(authHeader, " ")
		token := splitToken[1]
		if services.TOKEN == token {
			return ctx.Next()
		} else {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":       "token invalid",
				"OriginToken": services.TOKEN,
				"LocalToken":  authHeader,
			})
		}
	}
}
