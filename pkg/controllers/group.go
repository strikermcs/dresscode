package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satori/go.uuid"
	"github.com/strikermcs/dresscode/pkg/config"
	"github.com/strikermcs/dresscode/pkg/models"
	"github.com/strikermcs/dresscode/pkg/utils"
)

type Group struct {
	ID 		uuid.UUID
	Name 	string 	  `validate:"required,min=3,max=32"`
}

type GroupWithProducts struct {
	ID 			uuid.UUID
	Name 		string
	Products    []Product 	  
}

func CreateGroup(c *fiber.Ctx) error {
	g := new(Group)
	
	if err := c.BodyParser(g); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error()})
    }

	errors := utils.ValidateStruct(*g)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	group :=  models.Group{Name: g.Name}

	config.Database.Db.Create(&group)

	return c.Status(200).JSON(group.ID)
}

func GetAllGroups(c *fiber.Ctx) error {
	var groups []Group

	if err := config.Database.Db.Find(&groups).Error; err != nil {
		return c.Status(404).JSON("error: Record not found")
	}

	return c.Status(200).JSON(groups)
}

func GetGroupWithProducts(c *fiber.Ctx) error {
	g := new(Group)
	var products []Product

	if err := config.Database.Db.First(&g, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(404).JSON("Error: Group not Found")
	}
	
	if err := config.Database.Db.Find(&products, "group_id = ?", c.Params("id")).Error; err != nil {
		return c.Status(404).JSON("Error: Product not Found")
	}
	
	result := GroupWithProducts{ID: g.ID, Name: g.Name, Products: products}

	return c.Status(200).JSON(result)
}

func DeleteGroup(c *fiber.Ctx) error {
	if err := config.Database.Db.Unscoped().Delete(&models.Group{}, "id = ?", c.Params("id")).Error; err != nil {
		return c.JSON("error delete Groups")
	}
	return c.JSON(c.Params("id"))
}