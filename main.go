package main

import (
	"github.com/gofiber/fiber/v2"
	routes "github.com/skullblack-strak/sale-page-api/api"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")
	routes.SetupRouter(api)
	app.Listen(":3000")
}
