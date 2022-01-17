package main

import "github.com/gofiber/fiber/v2"

func MountGroupSubApp() *fiber.App{

	subApp := fiber.New();

	subApp.Get("/", func (c *fiber.Ctx) {
		c.SendString("This is version v2")
	})

	return subApp
}