package main



import "github.com/gofiber/fiber/v2"


func main() {
		app := fiber.New()

		app.Static("/html" , "./public/html")
		app.Get("/", func(c *fiber.Ctx) error {
				return c.SendString("HELLO WORLD")
		})



		app.Get("/params/path/:name", func(c *fiber.Ctx) error {
			return c.SendString("HELLO -> " + c.Params("name"))
		})


		app.Listen(":3000")

}