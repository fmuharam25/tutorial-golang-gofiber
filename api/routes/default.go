package routes

import (
	handlers "github.com/fmuharam25/tutorial-golang-gofiber/api/handlres"
	"github.com/gofiber/fiber/v2"
)

func DefaultRoute(app fiber.Router) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome API V1")
	})
	app.Post("/login", handlers.Login)
	app.Post("/logout", handlers.Logout)
}
