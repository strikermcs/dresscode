package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/strikermcs/dresscode/pkg/models"
)	


type Product struct {
	Id uint
	Name string
	ProductModel string
	Quantity uint
	Price float32
	Image string
	Code string
	OldPrice float32
}

func CreateProduct(c *fiber.Ctx) error {
	
	p := new(Product)
	
	if err := c.BodyParser(p); err != nil {
        return err
    }

	product := models.Product{Name: p.Name, ProductModel: p.ProductModel,
		 Quantity: p.Quantity, Price: p.Price, Image: p.Image, Code: p.Code, OldPrice: p.OldPrice}

	result := db.Create(&product)

	log.Println(result.Error)
	log.Println(result.RowsAffected)
	return c.JSON(product.ID)
}


func GetAllProducts(c *fiber.Ctx) error {

	var products []Product

	db.Find(&products)

	return c.JSON(products)
}


func GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")

	product := new(Product)
	
	db.First(&product, id)

	return c.JSON(product)
}