package controllers

import "github.com/gofiber/fiber/v2"

func GetEvents(ctx *fiber.Ctx) error {
	return ctx.SendString("GetEVents")
}

func AddEvent(ctx *fiber.Ctx) error {
	return ctx.SendString("AddEvent")
}

func DeleteEvent(ctx *fiber.Ctx) error {
	return ctx.SendString("DeleteEvent")
}

func UpdateEvent(ctx *fiber.Ctx) error {
	return ctx.SendString("UpdateEvent")
}
