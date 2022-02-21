package main

import (
	f "fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, any!")
}

func main() {
	f.Println("Welcome to the mothership")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/cake", func(c *fiber.Ctx) error {
		return c.SendString("Cake")
	})

	app.Use("*", hello)

	log.Fatal(app.Listen(":3000"))
}
