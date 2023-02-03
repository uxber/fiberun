package base

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func FiberEngine() *fiber.App {
	engine := fiber.New()

	engine.Use(cors.New())
	engine.Use(logger.New())
	engine.Use(recover.New())

	return engine
}
