package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"osa.synapsis.chalange/modules/user/model"
	"osa.synapsis.chalange/modules/user/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{userService}
}

func (u UserController) Register(c *fiber.Ctx) (err error) {
	var user model.User
	err = json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"status" : http.StatusText(400),
			"message" : err.Error(),
		})
	}

	err = u.userService.Register(user)
	if err != nil {
		return c.Status(500).JSON(map[string]interface{}{
			"status" : http.StatusText(500),
			"message" : err.Error(),
		})
	}

	return c.Status(201).JSON(map[string]interface{}{
			"status" : http.StatusText(201),
			"message" : "success",
		})
}

func (u UserController) Login(c *fiber.Ctx) (err error) {
	var user model.User
	err = json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"status" : http.StatusText(400),
			"message" : err.Error(),
		})
	}

	token, err := u.userService.Login(user)
	if err != nil {
		return c.Status(500).JSON(map[string]interface{}{
			"status" : http.StatusText(500),
			"message" : err.Error(),
		})
	}

	return c.Status(200).JSON(map[string]interface{}{
			"status" : http.StatusText(200),
			"message" : "success",
			"data" : map[string]any{
				"token" : token,
				"expired_in" : time.Minute * 5,
			},
		})
}
