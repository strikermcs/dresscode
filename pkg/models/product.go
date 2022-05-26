package models

import "github.com/satori/go.uuid"	

type Group struct {
	Base
	Name 			string		`json:"name"`
	Products 		[]Product	`gorm:"constraint:OnDelete:CASCADE;"`		
}

type Product struct {
	Base
	Name  		 	string		`json:"name"`
	ProductModel 	string
	Quantity  	 	uint64
	Price 			float32
	OldPrice 		float32
	Image 			string
	Sex 			string
	Code 			string
	GroupID 		uuid.UUID	
	Description 	Description `gorm:"constraint:OnDelete:CASCADE;"`
}

type Description struct {
	Base
	BrendName 		string
	ProductColor 	string
	Material 		string
	Sesson 			string
	Сomposition 	string
	Style 			string
	ProductSizes 	[]Size		`gorm:"many2many:description_sizes;constraint:OnDelete:CASCADE;"`
	Images 			[]Image		`gorm:"constraint:OnDelete:CASCADE;"`
	ProductID 		uuid.UUID
}

type Size struct {
	Base
	Name 					string		`json:"name"`	
	Size 					string
	DescriptionID 			uuid.UUID
}

type Image struct {
	Base
	Name 					string		`json:"name"`
	Path 					string
	DescriptionID 			uuid.UUID
}

