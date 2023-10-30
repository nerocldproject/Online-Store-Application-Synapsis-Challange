package repository

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"osa.synapsis.chalange/modules/cart/model"
	mp "osa.synapsis.chalange/modules/product/model"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return CartRepository{db}
}

func (c CartRepository) InsertCart(cart model.Cart) (err error) {
	tx := c.db.Begin()
	defer func(){
		if r := recover(); r != nil {
			tx.Rollback()
			err = r.(error)
		}
		
		if err != nil {
			return
		}
	}()

	var product mp.Product
	result := tx.Table("product").
				Where("product_id = ? and deleted_at is null", cart.ProductID).First(&product)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	if result.RowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("Maaf, Produk tidak tersedia")
	}
	
	var cartInDB model.Cart
	result = tx.Table("cart").
				Where("product_id = ? and user_id = ? and deleted_at is null and invoice_id is null", cart.ProductID, cart.UserID).
				First(&cartInDB)
	if result.Error != nil {
		if result.Error.Error() != "record not found" {
			tx.Rollback()
			return result.Error
		}
	}

	if cart.Quantity > product.Quantity {
		tx.Rollback()
		return fmt.Errorf("Maaf produk yang anda masukkan ke cart melebihi kapasitas")
	}

	if result.RowsAffected == 0 {
		result = tx.Table("cart").Omit("cart_id").Create(&cart)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}

	} else {
		result = tx.Table("cart").Where("product_id = ? and user_id = ? and deleted_at is null and invoice_id is null", cart.ProductID, cart.UserID).
				Update("quantity", cartInDB.Quantity+cart.Quantity)
	}

	result = tx.Table("product").Where("product_id = ? and deleted_at is null", product.ProductID).Update("quantity", product.Quantity-cart.Quantity)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}

	return tx.Commit().Error
}

func (c CartRepository) GetCartByUsername(username string) (carts []model.CartToResponse, err error) {
	err = c.db.Table("cart c").Select("p.name as product_name, c.quantity, p.price, (cast(c.quantity as bigint)*cast(p.price as bigint)) as total_price").
	Joins("join customer cu on cu.user_id = c.user_id").
	Joins("join product p on p.product_id = c.product_id").
	Where("cu.user_name = ? and c.deleted_at is null and c.invoice_id is null", username).Find(&carts).Error
	
	return
}

func (c CartRepository) DeleteCart() (err error) {
	tx := c.db.Begin()
	defer func(){
		if r := recover(); r != nil {
			tx.Rollback()
			err = r.(error)
		}
		
		if err != nil {
			return
		}
	}()
	
	var carts []model.Cart
	err = tx.Table("cart").Where("deleted_at is null and invoice_id is null").Find(&carts).Error
	if err != nil {
		tx.Rollback()
		return
	}

	fmt.Println(1)

	for _, cart := range carts {
		var product mp.Product

		err = tx.Table("product").Where("product_id = ? and deleted_at is null", cart.ProductID).Limit(1).Error
		if err != nil {
			tx.Rollback()
			return
		}

		
		err = tx.Table("product").Where("product_id = ? and deleted_at is null", product.ProductID).Update("quantity", product.Quantity+cart.Quantity).Error
		if err != nil {
			tx.Rollback()
			return
		}

		err = tx.Table("cart").Omit("updated_at").Where("cart_id = ? and deleted_at is null and invoice_id is null", cart.CartID).Update("deleted_at", time.Now()).Error
		if err != nil {
			tx.Rollback()
			return
		}
	}

	return tx.Commit().Error
}

func (c CartRepository) CreateInvoice(username string) (invoice model.InvoiceToRepsponse, err error) {
	tx := c.db.Begin()
	defer func(){
		if r := recover(); r != nil {
			tx.Rollback()
			err = r.(error)
		}

		if err != nil {
			return
		}
	}()

	invoiceID := uuid.New().String()

	inv := model.Invoice{
		InvoiceID: invoiceID,
	}

	err = tx.Table("invoice").Create(&inv).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Table("cart c").
		Joins("join costumer cu on cu.user_id = c.user_id").
		Where("cu.user_name = ? and c.deleted_at is null and c.invoice_id is null").Update("c.invoice_id", invoiceID).Error
	
	if err != nil {
		tx.Rollback()
		return
	}
	
	var carts []model.CartToResponse
	err = c.db.Table("cart c").Select("p.name, c.quantity, p.price, (cast(c.quantity as bigint)*cast(p.price as bigint)) as total_price").
	Joins("join customer cu on cu.user_id = c.user_id").
	Joins("join product p on p.product_id = c.product_id").
	Where("cu.user_name = ? and c.deleted_at is null and c.invoice_id = ?", username, invoiceID).Find(&carts).Error
	if err != nil {
		tx.Rollback()
		return
	}

	var totalPrice int64
	for _, cart := range carts {
		totalPrice += cart.TotalPrice
	}

	invoice = model.InvoiceToRepsponse{
		InvoiceID: invoiceID,
		CartToResponse: carts,
		TotalPrice: totalPrice,
	}
	return invoice, tx.Commit().Error
}

func (c CartRepository) GetInvoiceById(id string, username string) (invoice model.InvoiceToRepsponse, err error) {
	tx := c.db.Begin()
	defer func(){
		if r := recover(); r != nil {
			tx.Rollback()
			err = r.(error)
		}
		
		if err != nil {
			return
		}
	}()

	var carts []model.CartToResponse
	err = c.db.Table("cart c").Select("p.name, c.quantity, p.price, (cast(c.quantity as bigint)*cast(p.price as bigint)) as total_price").
	Joins("join customer cu on cu.user_id = c.user_id").
	Joins("join product p on p.product_id = c.product_id").
	Where("cu.user_name = ? and c.deleted_at is null and c.invoice_id = ?", username, id).Find(&carts).Error
	if err != nil {
		tx.Rollback()
		return
	}

	var totalPrice int64
	for _, cart := range carts {
		totalPrice += cart.TotalPrice
	}

	invoice = model.InvoiceToRepsponse{
		InvoiceID: id,
		CartToResponse: carts,
		TotalPrice: totalPrice,
	}
	return invoice, tx.Commit().Error
}