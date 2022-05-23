package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/strikermcs/dresscode/pkg/controllers"
)

func InitRoutes(r *fiber.App){
	r.Post("/product", controllers.CreateProduct)
	r.Get("/product", controllers.GetAllProducts)
	r.Get("/product/:id", controllers.GetProductById)
}