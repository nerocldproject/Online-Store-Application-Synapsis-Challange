package user

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"osa.synapsis.chalange/modules/user/controller"
	"osa.synapsis.chalange/modules/user/repository"
	"osa.synapsis.chalange/modules/user/service"
)

func NewUserHandler(g fiber.Router, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	g.Post("/register", userController.Register)
	g.Post("/login", userController.Login)
}