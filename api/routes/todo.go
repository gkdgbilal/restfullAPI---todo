package routes

import (
	"RestFullAPI-todo/api/utils"
	"RestFullAPI-todo/api/utils/middleware"
	"RestFullAPI-todo/configs/logg"
	"RestFullAPI-todo/pkg/auth"
	"RestFullAPI-todo/pkg/dto"
	"RestFullAPI-todo/pkg/todo"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
)

func TodoRouter(app fiber.Router, as auth.Service, service todo.Service) {
	//app.Post("/todos", middleware.Auth(as), createTodo(service))
	//app.Get("/todos", middleware.Auth(as), readTodos(service))
	//app.Get("/todos/:id", middleware.Auth(as), readTodo(service))
	//app.Put("/todos", middleware.Auth(as), updateTodo(service))
	//app.Put("/todos/done", middleware.Auth(as), completedTodo(service))
	//app.Delete("/todos/:id", middleware.Auth(as), removeTodo(service))
	app.Post("/todos", middleware.Auth(as, true), createTodo(service))
	app.Get("/todos", middleware.Auth(as, true), readTodos(service))
	app.Get("/todos/:id", middleware.Auth(as, true), readTodo(service))
	app.Put("/todos", middleware.Auth(as, true), updateTodo(service))
	app.Put("/todos/done", middleware.Auth(as, true), completedTodo(service))
	app.Delete("/todos/:id", middleware.Auth(as, true), removeTodo(service))
}

func completedTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody dto.TodoDTO
		err := c.BodyParser(&requestBody)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		result, dberr := service.Completed(&requestBody)
		return c.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}

func removeTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		//id := c.Params("id")
		//query := c.Query("hardDelete")
		//hardDelete, _ := utils.StrTo(query).Bool()
		//dberr := service.Delete(id, hardDelete)
		//if dberr != nil {
		//	logg.L.Error("err", zap.Error(dberr))
		//	return c.Status(http.StatusBadGateway).JSON(&fiber.Map{
		//		"status": http.StatusBadGateway,
		//		"error":  "Message: " + dberr.Error(),
		//	})
		//}
		//return c.Status(http.StatusNoContent).JSON("")
		id := c.Params("id")
		//query := c.Query("hardDelete")
		//hardDelete, _ := utils.StrTo(query).Bool()
		dberr := service.Delete(id)
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

func updateTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody dto.TodoDTO
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

func readTodo(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		reports, err := service.Read(id)
		if err != nil {
			return fiber.NewError(fiber.StatusConflict, err.Error())
		}
		return c.JSON(reports)
	}
}

func readTodos(service todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		pageable := utils.NewPagination()
		todos, err := service.Reads(pageable)
		if err != nil {
			return fiber.NewError(fiber.StatusConflict, err.Error())
		}
		return c.JSON(todos)
	}
}

func createTodo(s todo.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestBody := new(dto.TodoDTO)
		err := utils.ParseBodyAndValidate(c, requestBody)
		if err != nil {
			return err
		}
		todoResponse, dberr := s.Create(requestBody)
		if dberr != nil {
			return fiber.NewError(fiber.StatusConflict, dberr.Error())
		}
		return c.JSON(&todoResponse)
	}
}
