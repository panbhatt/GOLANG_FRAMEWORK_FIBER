package main

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)


func main() {
		app := fiber.New(fiber.Config{

			ServerHeader :"Express",
			AppName : "Test App v1.0.1",
		})

		app.Static("/html" , "./public/html", fiber.Static {
			Index : "index.html",
			CacheDuration : 10 * time.Second,
		})
		app.Get("/", func(c *fiber.Ctx) error {
				return c.SendString("HELLO WORLD")
		})



		app.Get("/params/path/:name", func(c *fiber.Ctx) error {
			return c.SendString("HELLO -> " + c.Params("name"))
		})

		app.Get("/error", func(c *fiber.Ctx) error {
			return fiber.NewError(http.StatusBadRequest, "Custom Error message")
		})

		// API Related Routes.
		app.Use("/api/*", func (c *fiber.Ctx) error {
			fmt.Println("MIddleware Invokced")
			return c.Next()
		})

		app.Get("/api/v1", func(c *fiber.Ctx) error {
			return c.SendString("Returning the APi Version v1")
		})

		// Mount SubAPp
		app.Mount("subapp", GetSubApp())
		Mount()


		app.Listen(":3000")

}

func GetSubApp1() *fiber.App {
	subAp := fiber.New()
	subAp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Sending it from sub app")
	})
	return subAp
}