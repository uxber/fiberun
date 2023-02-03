package service

import (
	"fiberun/internal/handler"
	"fiberun/internal/model"
	"github.com/gofiber/fiber/v2"
)

type UserService struct{}

func (UserService) SignIn(ctx *fiber.Ctx) error {
	user := model.User{}
	userCtl := handler.UserHandler{}
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(Response{Code: 400, Message: err.Error()})
	}
	if err := userCtl.ValidateParams(&user); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(Response{Code: 400, Message: "Authentication failed"})
	}

	findUser, ok := userCtl.FindOne(model.User{Username: user.Username})
	err := userCtl.CompareHash(&findUser, user.Password)
	if !ok || err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(Response{Code: 400, Message: "Authentication failed"})
	}

	token, err := userCtl.GenerateToken(&user)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(Response{Code: 200, Message: "Success", Data: fiber.Map{"token": token}})
}

func (UserService) SignUp(ctx *fiber.Ctx) error {
	user := model.User{}
	userCtl := handler.UserHandler{}
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(Response{Code: 400, Message: err.Error()})
	}
	if err := userCtl.ValidateParams(&user); err != nil {
		return ctx.JSON(Response{Code: 400, Message: err.Error()})
	}

	if _, exist := userCtl.FindOne(model.User{Username: user.Username}); exist {
		return ctx.Status(fiber.StatusBadRequest).JSON(Response{Code: 400, Message: "User already exists"})
	}
	if err := userCtl.EncryptHash(&user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(Response{Code: 500, Message: err.Error()})
	}
	if err := userCtl.CreateUser(&user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(Response{Code: 500, Message: err.Error()})
	}

	return ctx.JSON(Response{Code: 200, Message: "Success", Data: user})
}
