package service

import (
	"osa.synapsis.chalange/modules/product/model"
	"osa.synapsis.chalange/modules/product/repository"
)

type ProductService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return ProductService{productRepo}
}

func (p ProductService) InsertProduct(product model.Product) (err error) {
	err = p.productRepo.InsertProduct(product)
	return
}

func (p ProductService) UpdateProduct(product model.Product) (err error) {
	err = p.productRepo.UpdateProduct(product)
	return
}

func (p ProductService) DeleteProduct(id int) (err error) {
	err = p.productRepo.DeleteProduct(id)
	return
}

func (p ProductService) GetProductById(id int) (user model.Product, err error) {
	user, err = p.productRepo.GetProductById(id)
	return
}

func (p ProductService) GetAllProduct() (users []model.Product, err error) {
	users, err = p.productRepo.GetAllProduct()
	return
}