package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Title       string
	Description string
	Price       float64
	ImageUrl    string
	SellerID    uint
	Seller      User
	PurchaserID *uint
	Purchaser   User
}
