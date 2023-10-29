package product

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"osa.synapsis.chalange/modules/product/controller"
	"osa.synapsis.chalange/modules/product/repository"
	"osa.synapsis.chalange/modules/product/service"
	"osa.synapsis.chalange/modules/user/model"
	"osa.synapsis.chalange/utils/middleware"
)

func NewProductHandler(g fiber.Router, db *gorm.DB) {
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)
	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: viper.GetString("SECRET_KEY")},
		Claims: &model.Claims{},
	})

	g.Post("/", jwtMiddleware, middleware.Authorization, productController.InsertProduct)
	g.Put("/:id", jwtMiddleware, middleware.Authorization, productController.UpdateProduct)
	g.Delete("/:id", jwtMiddleware, middleware.Authorization, productController.DeleteProduct)
	g.Get("/:id", productController.GetProductById)
	g.Get("/", productController.GetAllProduct)
}