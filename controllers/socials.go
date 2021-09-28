package controllers

import "github.com/gofiber/fiber/v2"

func GetSocials(ctx *fiber.Ctx) error {
	return ctx.SendString("GetSocials")
}

func AddSocial(ctx *fiber.Ctx) error {
	return ctx.SendString("UpdateSocial")
}

func DeleteSocial(ctx *fiber.Ctx) error {
	return ctx.SendString("UpdateSocial")
}

func UpdateSocial(ctx *fiber.Ctx) error {
	return ctx.SendString("UpdateSocial")
}
