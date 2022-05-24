package models

type Product struct {
	Base
	Name  		 	string	
	ProductModel 	string
	Quantity  	 	uint64
	Price 			float32
	OldPrice 		float32
	Image 			string
	Code 			string
	Description 	ProductDescription 
}

type ProductDescription struct {
	Base
	BrendName 		string
	ProductColor 	string
	Material 		string
	Sex 			string
	Sesson 			string
	Сomposition 	string
	Style 			string
	ProductSizes 	[]ProductSize
	Images 			[]ProductImage
	ProductID 		uint
}

type ProductSize struct {
	Base
	Name 					string
	Size 					string
	ProductDescriptionID 	uint
}

type ProductImage struct {
	Base
	Name 					string
	Path 					string
	ProductDescriptionID 	uint
}

