package routes

import (
	"RestFullAPI-todo/api/utils"
	"RestFullAPI-todo/api/utils/middleware"
	"RestFullAPI-todo/configs/logg"
	"RestFullAPI-todo/pkg/auth"
	"RestFullAPI-todo/pkg/dto"
	"RestFullAPI-todo/pkg/user"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func UserRouter(app fiber.Router, as auth.Service, service user.Service) {
	app.Get("/users", middleware.Auth(as), readUsers(service))
	app.Get("/users/:id", middleware.Auth(as, true), readUser(service))
	app.Get("/users/u/:username", middleware.Auth(as, true), readUserByUsername(service))
	app.Put("/users", middleware.Auth(as), updateUsers(service))
	app.Delete("/users/:id", middleware.Auth(as), removeUsers(service))
}

func readUsers(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		time.Sleep(time.Second * 2)
		users, err := service.Reads()
		if err != nil {
			return fiber.NewError(fiber.StatusConflict, err.Error())
		}
		return c.JSON(users)
	}
}

func readUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		users, err := service.Read(id)
		if err != nil {
			return fiber.NewError(fiber.StatusConflict, err.Error())
		}
		return c.JSON(users)
	}
}

func readUserByUsername(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Params("username")
		users, err := service.ReadByUsername(username)
		if err != nil {
			return fiber.NewError(fiber.StatusConflict, err.Error())
		}
		return c.JSON(users)
	}
}

func updateUsers(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody dto.UserDTO
		err := c.BodyParser(&requestBody)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		result, dberr := service.Update(&requestBody)
		return c.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}

func removeUsers(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		query := c.Query("hardDelete")

		hardDelete, _ := utils.StrTo(query).Bool()

		dberr := service.Delete(id, hardDelete)

		if dberr != nil {
			logg.L.Error("err", zap.Error(dberr))
			return c.Status(http.StatusBadGateway).JSON(&fiber.Map{
				"status": http.StatusBadGateway,
				"error":  "Message: " + dberr.Error(),
			})
		}
		return c.Status(http.StatusNoContent).JSON("")
	}
}
