package middlewares

import "github.com/gofiber/fiber/v2"

func CheckAuth(ctx *fiber.Ctx) error {
	return ctx.SendString("Auth")
}
