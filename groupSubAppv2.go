package main

import "github.com/gofiber/fiber/v2"

func MountGroupSubApp() *fiber.App{

	subApp := fiber.New();

	subApp.Get("/", func(c *fiber.Ctx) error{
		c.SendString("This is version v1")
	})

	return subApp
}