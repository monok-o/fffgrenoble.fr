package controllers

import "github.com/gofiber/fiber/v2"

func GetLinks(ctx *fiber.Ctx) error {
	return ctx.SendString("GetLinks")
}

func AddLink(ctx *fiber.Ctx) error {
	return ctx.SendString("AddLink")
}

func DeleteLink(ctx *fiber.Ctx) error {
	return ctx.SendString("DeleteLink")
}

func UpdateLink(ctx *fiber.Ctx) error {
	return ctx.SendString("UpdateLink")
}
