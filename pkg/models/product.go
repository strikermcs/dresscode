package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name string
	ProductModel string
	Quantity uint
	Price float32
	OldPrice float32
	Description ProductDescription 
}

type ProductDescription struct {
	gorm.Model
	BrendName string
	ProductColor string
	Material string
	Sex string
	Sesson string
	Сomposition string
	Style string
	ProductID uint
}

type ProductSize struct {
	gorm.Model
	Name string
	
}