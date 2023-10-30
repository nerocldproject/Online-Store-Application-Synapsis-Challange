package repository

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"osa.synapsis.chalange/modules/product/model"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository (db *gorm.DB) ProductRepository {
	return ProductRepository{db}
}

func (p ProductRepository) InsertProduct(product model.Product) (err error) {
	result := p.db.Table("product").Omit("product_id").Create(&product)
	if result.Error != nil {
		return result.Error
	}

	return
}

func (p ProductRepository) UpdateProduct(product model.Product) (err error) {
	result := p.db.Table("product").Omit("created_at", "deleted_at").Where("product_id = ? and deleted_at is null").Updates(&product)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Maaf, produk tidak ada di database")
	}

	return
}

func (p ProductRepository) DeleteProduct(id int) (err error) {
	result := p.db.Table("product").Omit("updated_at").Where("product_id = ? and deleted_at is null", id).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Maaf, produk tidak ada di database")
	}

	return
}

func (p ProductRepository) GetProductById(id int) (product model.Product, err error) {
	result := p.db.Table("product").Where("product_id = ? and deleted_at is null", id).First(&product)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	
	return
}

func (p ProductRepository) GetProductByCategory(cat string) (products []model.Product, err error) {
	result := p.db.Table("product").Where("category_id = ? and deleted_at is null", cat).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	
	return
}

func (p ProductRepository) GetAllProduct() (products []model.Product, err error) {
	result := p.db.Table("product").Where("deleted_at is null").Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return
}