package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"osa.synapsis.chalange/modules/product/model"
	"osa.synapsis.chalange/modules/product/service"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return ProductController{productService}
}

func (p ProductController) InsertProduct(c *fiber.Ctx) (err error) {
	var product model.Product
	err = json.Unmarshal(c.Body(), &product)
	if err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"status" : http.StatusText(400),
			"message" : err.Error(),
		})
	}

	err = p.productService.InsertProduct(product)
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

func (p ProductController) UpdateProduct(c *fiber.Ctx) (err error) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"status" : http.StatusText(400),
			"message" : err.Error(),
		})
	}
	
	var product model.Product
	err = json.Unmarshal(c.Body(), &product)
	if err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"status" : http.StatusText(400),
			"message" : err.Error(),
		})
	}

	product.ProductID = id
	err = p.productService.UpdateProduct(product)
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

func (p ProductController) DeleteProduct(c *fiber.Ctx) (err error) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"status" : http.StatusText(400),
			"message" : err.Error(),
		})
	}

	err = p.productService.DeleteProduct(id)
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

func (p ProductController) GetProductById(c *fiber.Ctx) (err error) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"status" : http.StatusText(400),
			"message" : err.Error(),
		})
	}

	product, err := p.productService.GetProductById(id)
	if err != nil {
		return c.Status(500).JSON(map[string]interface{}{
			"status" : http.StatusText(500),
			"message" : err.Error(),
		})
	}

	return c.Status(200).JSON(map[string]interface{}{
			"status" : http.StatusText(200),
			"message" : "success",
			"data" : product,
		})
}

func (p ProductController) GetAllProduct(c *fiber.Ctx) (err error) {
	var products []model.Product
	cat := c.Query("cat")
	if cat == "" {
		products, err = p.productService.GetAllProduct()
	} else {
		products, err = p.productService.GetProductByCategory(cat)
	}
	
	if err != nil {
		return c.Status(500).JSON(map[string]interface{}{
			"status" : http.StatusText(500),
			"message" : err.Error(),
		})
	}

	return c.Status(200).JSON(map[string]interface{}{
			"status" : http.StatusText(200),
			"message" : "success",
			"data" : products,
		})
}