package router

import (
	"fiberun/internal/service"
	"fiberun/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(engine *fiber.App) {

	// No authorization required
	engine.Post("/signin", service.UserGroup.SignIn)
	engine.Post("/signup", service.UserGroup.SignUp)

	// Restricted Routes
	api := engine.Group("/api").Use(middleware.JwtAuth())
	{
		api.Get("/test", func(ctx *fiber.Ctx) error {
			return ctx.SendString("success")
		})
	}

	// 404 Handler
	engine.Use(middleware.NotFound)
}
