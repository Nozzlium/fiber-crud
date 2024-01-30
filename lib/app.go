package lib

import (
	"database/sql"
	"fiber-crud/controller"
	"fiber-crud/repository"
	"fiber-crud/service"

	"github.com/gofiber/fiber/v2"
)

func NewApp(db *sql.DB) *fiber.App {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(
		db,
		userRepository,
	)
	userController := controller.NewUserController(userService)

	app := fiber.New()

	user := app.Group("/user")
	user.Post("", userController.Create)
	user.Get("", userController.Get)
	user.Get("/:id", userController.GetById)
	user.Put("/:id", userController.Update)
	user.Delete("/:id", userController.Delete)

	return app
}
