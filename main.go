package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"

	"time"
)

type Person struct {
	Name string `json:"name"`
	Age uint8 `json:"age"`
}

func main() {
	app := fiber.New(fiber.Config{

		ServerHeader: "Express",
		AppName:      "Test App v1.0.1",
	})

	app.Static("/html", "./public/html", fiber.Static{
		Index:         "index.html",
		CacheDuration: 10 * time.Second,
	}).Name("Getting the HTML Page")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("HELLO WORLD")
	})

	app.Get("/params/path/:name", func(c *fiber.Ctx) error {
		return c.SendString("HELLO -> " + c.Params("name"))
	})

	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(http.StatusBadRequest, "Custom Error message")
	})

	// Mount SubAPp
	app.Mount("subapp", GetSubApp())
	Mount()

	// Group Sub APP
	app.Mount("/api/v1", MountGroupSubAppV1())
	app.Mount("/api/v2", MountGroupSubAppV2())

	// API Related Routes.
	app.Use("/api/*", func(c *fiber.Ctx) error {
		fmt.Println("MIddleware Invokced")
		return c.Next()
	})

	app.Get("/api/p1", func(c *fiber.Ctx) error {
		c.Append("Link", "http:google.com")
		return c.SendString("Returning the APi Version v1")
	})

	app.Get("/stack", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})

	app.Post("/adduser", func(c *fiber.Ctx) error {
		p := new(Person)

		if err := c.BodyParser(p) ; err !=nil {
			return err
		}
			fmt.Println(p.Name)
			return c.JSON(p)


	})

	app.Get("/json", func(c *fiber.Ctx) error {
		p := Person{
			Name: "Pankaj", Age: 20,
		}
		return c.JSON(p)
	})

	app.Get("/401", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	app.Get("/readme", func(c *fiber.Ctx) error {
		return c.SendFile("./README.md", false)
	})

	app.Get("/jsonmap", func(c *fiber.Ctx) error {

		c.Links(
			"http://api.example.com/users?page=2", "next",
			"http://api.example.com/users?page=5", "last",
		)
		return c.JSON(fiber.Map{
			"name" : "Pankaj Bhatt",
			"age" : 20,
		})
	})

	fmt.Printf("Total Number of handers = %d", app.HandlersCount())
	// THis will give a list of all the Handlers the app has.
	fmt.Println("\nMarshalling Indent of Handlers()")
	data, _ := json.MarshalIndent(app.Stack(),"", " ")
	fmt.Println(string(data))

	app.Listen(":3000")

}

func GetSubApp1() *fiber.App {
	subAp := fiber.New()
	subAp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Sending it from sub app")
	})
	return subAp
}
