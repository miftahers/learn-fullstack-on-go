package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	// Load Templates
	engine := html.New("./views", ".tmpl")

	// Creating fiber app
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)

	//Configure app
	app.Static("/", "./public")

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"name": "Miftah Firdaus",
		})
	})

	// Serve on :3000
	app.Listen(":3000")
}
