package main

import "github.com/gofiber/fiber/v2"

func MountGroupSubAppV1() *fiber.App {

	subApp := fiber.New()

	subApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is version v1")
	}).Name("V1 Default get page")

	return subApp
}
