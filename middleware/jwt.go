package middleware

import (
	"fiberun/module/conf"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func JwtAuth() func(*fiber.Ctx) error {
	var secret = []byte(conf.Get("jwt.secret"))
	return jwtware.New(jwtware.Config{
		SigningKey: secret,
	})
}
