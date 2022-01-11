package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"net/http"
)


func main() {
		app := fiber.New(fiber.Config{

			ServerHeader :"Express",
			AppName : "Test App v1.0.1",
		})

		app.Static("/html" , "./public/html")
		app.Get("/", func(c *fiber.Ctx) error {
				return c.SendString("HELLO WORLD")
		})



		app.Get("/params/path/:name", func(c *fiber.Ctx) error {
			return c.SendString("HELLO -> " + c.Params("name"))
		})

		app.Get("/error", func(c *fiber.Ctx) error {
			return fiber.NewError(http.StatusBadRequest, "Custom Error message")
		})


		app.Listen(":3000")

}