package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/strikermcs/dresscode/pkg/controllers"
)

func ProductRoutesInit(r *fiber.App){
	r.Post("/product", controllers.CreateProduct)
	r.Get("/product", controllers.GetAllProducts)
	r.Get("/product/:id", controllers.GetProductById)
	r.Put("/product/:id-:quantity", controllers.PutBuyProduct)
	r.Put("/product", controllers.UpdateProduct)
	r.Delete("/product/:id", controllers.DeleteProduct)
}
	
