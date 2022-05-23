package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  		 string	
	ProductModel string
	Quantity  	 uint
	Price float32
	OldPrice float32
	Image string
	Code string
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
	ProductSizes []ProductSize
	Images []ProductImage
	ProductID uint
}

type ProductSize struct {
	gorm.Model
	Name string
	Size string
	ProductDescriptionID uint
}

type ProductImage struct {
	gorm.Model
	Name string
	Path string
	ProductDescriptionID uint
}

