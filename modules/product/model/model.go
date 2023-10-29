package model

import "time"

type Product struct {
	ProductID  int    `json:"product_id"`
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
	Price      int    `json:"price"`
	Quantity int `json:"quantity"`
	CreatedAt  *time.Time `json:"-"`
	UpdateAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}