package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"osa.synapsis.chalange/modules/cart"
	"osa.synapsis.chalange/modules/product"
	"osa.synapsis.chalange/modules/user"
)

func NewRoutes(app *fiber.App, db *gorm.DB) {
	v1 := app.Group("/v1")
	user.NewUserHandler(v1.Group("/user"), db)
	product.NewProductHandler(v1.Group("/product"), db)
	cart.NewCartHandler(v1.Group("/cart"), db)
}
