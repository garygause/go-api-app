package models

import "time"

type Product struct {
	ID int
	Title string `binding:"required"`
	Description string
	Price float64
	CreatedAt time.Time
	Status string
	StoreID int
}

var products = []Product{}

func (p Product) Save() {
	// TODO: save to db
	products = append(products, p)
}

func GetAllProducts() []Product {
	return products
}