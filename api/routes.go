package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(router fiber.Router) {
	about := router.Group("/about")

	// start middleware
	// Match request starting with /api
	// about.Use("/api", func(c *fiber.Ctx) error {
	// 	return c.Next()
	// })
	// // Match requests starting with /api or /home (multiple-prefix support)
	// about.Use([]string{"/api", "/home"}, func(c *fiber.Ctx) error {
	// 	return c.Next()
	// })
	// end middleware

	about.Get("/", func(c *fiber.Ctx) error { return c.SendString("path => " + c.Path()) })

}
