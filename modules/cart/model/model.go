package model

import "time"

type Cart struct {
	CartID    int        `json:"cart_id"`
	ProductID int        `json:"product_id"`
	UserID    int        `json:"user_id"`
	Quantity  int        `json:"quantity"`
	InvoiceID *string    `json:"-"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type CartToResponse struct {
	ProductName string `json:"product_name"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
	TotalPrice int64 `json:"total_price"`
}

type Invoice struct {
	InvoiceID string `json:"invoice_id"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type InvoiceToRepsponse struct {
	InvoiceID string `json:"invoice_id"`
	CartToResponse []CartToResponse `json:"carts"`
	TotalPrice int64 `json:"total_price"`
}