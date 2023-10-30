package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"osa.synapsis.chalange/modules/cart/model"
	"osa.synapsis.chalange/modules/cart/service"
)

type CartController struct {
	cartService service.CartService
}

func NewCartController(cartService service.CartService) CartController {
	return CartController{cartService}
}

func (cc CartController) InsertCart(c *fiber.Ctx) error {
	var cart model.Cart
	err := json.Unmarshal(c.Body(), &cart)
	if err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"status" : http.StatusText(400),
			"message" : err.Error(),
		})
	}

	err = cc.cartService.InsertCart(cart)
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

func (cc CartController) GetCartByUsername(c *fiber.Ctx) error {
	username := c.Locals("user_name").(string)
	if username == "" {
		return c.Status(400).JSON(map[string]interface{}{
			"status" : http.StatusText(400),
			"message" : "username tidak boleh kosong",
		})
	}

	carts, err := cc.cartService.GetCartByUsername(username)
	if err != nil {
		return c.Status(500).JSON(map[string]interface{}{
			"status" : http.StatusText(500),
			"message" : err.Error(),
		})
	}

	return c.Status(200).JSON(map[string]interface{}{
		"status" : http.StatusText(200),
		"message" : "success",
		"data" : carts,
	})
}

func (cc CartController) DeleteCart(c *fiber.Ctx) error {
	err := cc.cartService.DeleteCart()
	if err != nil {
		return c.Status(500).JSON(map[string]interface{}{
			"status" : http.StatusText(500),
			"message" : err.Error(),
		})
	}

	return c.Status(200).JSON(map[string]interface{}{
		"status" : http.StatusText(200),
		"message" : "success",
	})
}

func (cc CartController) CheckOut(c *fiber.Ctx) error {
	username := c.Locals("user_name").(string)
	if username == "" {
		return c.Status(400).JSON(map[string]interface{}{
			"status" : http.StatusText(400),
			"message" : "username tidak boleh kosong",
		})
	}

	invoice, err := cc.cartService.CreateInvoice(username)
	if err != nil {
		return c.Status(500).JSON(map[string]interface{}{
			"status" : http.StatusText(500),
			"message" : err.Error(),
		})
	}

	return c.Status(200).JSON(map[string]interface{}{
		"status" : http.StatusText(200),
		"message" : "success",
		"data" : invoice,
	})
}

func (cc CartController) GetInvoiceById(c *fiber.Ctx) error {
	username := c.Locals("user_name").(string)
	if username == "" {
		return c.Status(400).JSON(map[string]interface{}{
			"status" : http.StatusText(400),
			"message" : "username tidak boleh kosong",
		})
	}

	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(map[string]interface{}{
			"status" : http.StatusText(400),
			"message" : "id invoice tidak boleh kosong",
		})
	}

	invoice, err := cc.cartService.GetInvoiceById(id, username)
	if err != nil {
		return c.Status(500).JSON(map[string]interface{}{
			"status" : http.StatusText(500),
			"message" : err.Error(),
		})
	}

	return c.Status(200).JSON(map[string]interface{}{
		"status" : http.StatusText(200),
		"message" : "success",
		"data" : invoice,
	})
}