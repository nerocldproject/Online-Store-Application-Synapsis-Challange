package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"osa.synapsis.chalange/modules/user/model"
)

func Authorization(c *fiber.Ctx) (err error) {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(500).JSON(map[string]interface{}{
			"status" : http.StatusText(500),
			"message" : "Error parsing token",
		})
	}

	claims := token.Claims.(*model.Claims)
	if !ok {
		return c.Status(500).JSON(map[string]interface{}{
			"status" : http.StatusText(500),
			"message" : "Error parsing claims from token",
		})
	}

	c.Locals("user_name", claims.Username)
	return c.Next()
}