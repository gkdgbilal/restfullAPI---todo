package main

import (
	"RestFullAPI-todo/api/database"
	"RestFullAPI-todo/api/routes"
	"RestFullAPI-todo/api/utils"
	"RestFullAPI-todo/configs"
	"RestFullAPI-todo/configs/logg"
	"RestFullAPI-todo/pkg/auth"
	"RestFullAPI-todo/pkg/todo"
	"RestFullAPI-todo/pkg/user"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
	"os"
)

func main() {
	configs.Init()
	logg.Init()
	//fs.Init()

	logg.L.Info("Application:\n",
		zap.String("application", configs.C.App.Name),
		zap.String("environment", configs.C.Env),
		zap.String("port", configs.C.App.Port),
	)

	db, err := database.Connect()
	if err != nil {
		logg.L.Fatal("Database Connection Message $s", zap.Error(err))
	}
	logg.L.Info("Database connection success!")

	userCollection := db.Collection(database.Users)
	todoCollection := db.Collection(database.Todos)

	userRepository := user.NewRepository(userCollection)
	todoRepository := todo.NewRepository(todoCollection)

	userService := user.NewService(userRepository)
	authService := auth.NewService(userRepository)
	todoService := todo.NewService(todoRepository)

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		// %s |%s %3d %s| %7v | %15s |%s %-7s %s
		Format: "[${time}] | ${yellow}${status}${reset} | ${latency} 	| ${method} ${path} ${queryParams}\n",
		TimeFormat: "15:04:05",
		TimeZone:   "Local",
		Output:     os.Stderr,
	}))

	v1 := app.Group("v1")
	routes.AuthRouter(v1, authService)
	routes.UserRouter(v1, authService, userService)
	routes.TodoRouter(v1, authService, todoService)

	logg.L.Fatal("Started", zap.Error(app.Listen(fmt.Sprintf(":%s", configs.C.App.Port))))
}
