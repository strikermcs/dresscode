package controllers

import (
	"strconv"
	"github.com/satori/go.uuid"
	"github.com/gofiber/fiber/v2"
	"github.com/strikermcs/dresscode/pkg/config"
	"github.com/strikermcs/dresscode/pkg/models"
)	


type Product struct {
	ID uuid.UUID
	Name string
	ProductModel string
	Quantity uint64
	Price float32
	Image string
	Code string
	OldPrice float32
}

func CreateProduct(c *fiber.Ctx) error {
	
	p := new(Product)
	
	if err := c.BodyParser(p); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error()})
    }

	product := models.Product{Name: p.Name, ProductModel: p.ProductModel,
		 Quantity: p.Quantity, Price: p.Price, Image: p.Image, Code: p.Code, OldPrice: p.OldPrice}

	config.Database.Db.Create(&product)

	return c.Status(200).JSON(product.ID)
}


func GetAllProducts(c *fiber.Ctx) error {

	var products []Product

	if err := config.Database.Db.Find(&products).Error; err != nil {
		return c.Status(404).JSON("error: Record not found")
	}

	return c.Status(200).JSON(products)
}


func GetProductById(c *fiber.Ctx) error {

	product := new(Product)
	
	if err := config.Database.Db.First(&product, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(404).JSON("Error: Product not Found")
	}

	return c.Status(200).JSON(product)
}

func PutBuyProduct(c *fiber.Ctx) error {

	var product models.Product

	if err := config.Database.Db.First(&product, "id = ?", c.Params("id")).Error; err != nil {
		return c.JSON("Error: Product not Found")
	}

	q, err := strconv.ParseUint(c.Params("quantity"), 10, 64)

	if err != nil {
		return c.JSON("Error with product quantity")
	}
 
	if product.Quantity < q {
		return c.JSON("Not enough products")
	} 

	product.Quantity = product.Quantity - q
	
	if err := config.Database.Db.Save(&product).Error; err != nil {
		return c.JSON("Error Save Product")
	}
	return c.JSON(product.Quantity)	

}

func UpdateProduct(c *fiber.Ctx) error {
	p := new(Product)
	
	if err := c.BodyParser(p); err != nil {
        return err
    }

	var product models.Product
	if err := config.Database.Db.First(&product, "id = ?", p.ID).Error; err != nil {
		return c.JSON("Error product not found")
	}
    
	product.Name = p.Name
	product.Quantity = p.Quantity
	product.Code = p.Code
	product.Image = p.Image
	product.ProductModel = p.ProductModel
	product.Price = p.Price
	product.OldPrice = p.OldPrice

	if err := config.Database.Db.Save(&product).Error; err != nil {
		return c.JSON("error to update product")
	}

	return c.JSON(product)

} 

func DeleteProduct(c *fiber.Ctx) error {
	if err := config.Database.Db.Unscoped().Delete(&models.Product{}, "id = ?", c.Params("id")).Error; err != nil {
		return c.JSON("error delete Product")
	}
	return c.JSON(c.Params("id"))
}