package models

import "github.com/satori/go.uuid"	

type Group struct {
	Base
	Name 			string
	Products 		[]Product			
}

type Product struct {
	Base
	Name  		 	string	
	ProductModel 	string
	Quantity  	 	uint64
	Price 			float32
	OldPrice 		float32
	Image 			string
	Sex 			string
	Code 			string
	GroupID 		uuid.UUID	
	Description 	Description 
}

type Description struct {
	Base
	BrendName 		string
	ProductColor 	string
	Material 		string
	Sesson 			string
	Сomposition 	string
	Style 			string
	ProductSizes 	[]Size	`gorm:"many2many:description_sizes;"`
	Images 			[]Image
	ProductID 		uuid.UUID
}

type Size struct {
	Base
	Name 					string
	Size 					string
	DescriptionID 			uuid.UUID
}

type Image struct {
	Base
	Name 					string
	Path 					string
	DescriptionID 			uuid.UUID
}

