package main

import (
	"fibergo/database"

	"github.com/gofiber/fiber"
)

func main() {
	db.Mydb() // <-- Connect DB

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.SendString("Hai Dari Fiber")
	})

	app.Listen(":3000")

}
