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
	return p.productRepo.InsertProduct(product)
}

func (p ProductService) UpdateProduct(product model.Product) (err error) {
	return p.productRepo.UpdateProduct(product)
}

func (p ProductService) DeleteProduct(id int) (err error) {
	return p.productRepo.DeleteProduct(id)
}

func (p ProductService) GetProductById(id int) (product model.Product, err error) {
	return p.productRepo.GetProductById(id)
}

func (p ProductService) GetProductByCategory(cat string) (products []model.Product, err error) {
	return p.productRepo.GetProductByCategory(cat)
}

func (p ProductService) GetAllProduct() (products []model.Product, err error) {
	return p.productRepo.GetAllProduct()
}