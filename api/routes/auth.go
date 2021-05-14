package routes

import (
	"RestFullAPI-todo/api/utils"
	"RestFullAPI-todo/pkg/auth"
	"RestFullAPI-todo/pkg/dto"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, service auth.Service) {
	app.Post("/sign-up", signUp(service))
	app.Post("/log-in", logIn(service))
}

func logIn(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestBody := new(dto.LoginDTO)
		err := utils.ParseBodyAndValidate(c, requestBody)
		if err != nil {
			return err
		}
		authResponse, err := service.Login(requestBody)
		if err != nil {
			return err
		}
		return c.JSON(&authResponse)
	}
}

func signUp(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestBody := new(dto.RegisterDTO)
		err := utils.ParseBodyAndValidate(c, requestBody)
		if err != nil {
			return err
		}
		authResponse, err := service.SignUp(requestBody)
		if err != nil {
			return err
		}
		return c.JSON(&authResponse)
	}
}
