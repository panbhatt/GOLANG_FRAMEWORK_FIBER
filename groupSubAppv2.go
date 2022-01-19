package main

import "github.com/gofiber/fiber/v2"

func MountGroupSubAppV2() *fiber.App {

	subApp := fiber.New()

	subApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is version V2")
	})

	return subApp
}
