package service

import (
	"osa.synapsis.chalange/modules/cart/model"
	"osa.synapsis.chalange/modules/cart/repository"
)

type CartService struct {
	cartRepo repository.CartRepository
}

func NewCartService(cartRepo repository.CartRepository) CartService {
	return CartService{cartRepo}
}

func (c CartService) InsertCart(cart model.Cart) (err error) {
	return c.cartRepo.InsertCart(cart)
}

func (c CartService) GetCartByUsername(username string) (carts []model.CartToResponse, err error) {
	return c.cartRepo.GetCartByUsername(username)
}

func (c CartService) DeleteCart() (err error) {
	return c.cartRepo.DeleteCart()
}

func (c CartService) CreateInvoice(username string) (invoice model.InvoiceToRepsponse, err error) {
	return c.cartRepo.CreateInvoice(username)
}

func (c CartService) GetInvoiceById(id, username string) (invoice model.InvoiceToRepsponse, err error) {
	return c.cartRepo.GetInvoiceById(id, username)
}