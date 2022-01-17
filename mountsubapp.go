package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetSubApp() *fiber.App {
	subAp := fiber.New()
	subAp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Sending it from sub app")
	})
	return subAp
}

func Mount(){
	fmt.Println("IN Mount function")
}
