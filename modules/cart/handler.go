package cart

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"osa.synapsis.chalange/modules/cart/controller"
	"osa.synapsis.chalange/modules/cart/repository"
	"osa.synapsis.chalange/modules/cart/service"
	"osa.synapsis.chalange/modules/user/model"
	"osa.synapsis.chalange/utils/middleware"
)

func NewCartHandler(g fiber.Router, db *gorm.DB) {
	cartRepo := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepo)
	cartController := controller.NewCartController(cartService)
	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(viper.GetString("SECRET_KEY"))},
		Claims: &model.Claims{},
	})

	g.Post("/", jwtMiddleware, middleware.Authorization, cartController.InsertCart)
	g.Get("/", jwtMiddleware, middleware.Authorization, cartController.GetCartByUsername)
	g.Delete("/", jwtMiddleware, middleware.Authorization, cartController.DeleteCart)
	g.Post("/checkout", jwtMiddleware, middleware.Authorization, cartController.CheckOut)
	g.Get("/:id", jwtMiddleware, middleware.Authorization, cartController.GetInvoiceById)
}