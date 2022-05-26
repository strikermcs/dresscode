package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/strikermcs/dresscode/pkg/controllers"
)

func GroupsRouterInit(r *fiber.App){
	r.Post("/groups", controllers.CreateGroup)
	r.Get("/groups", controllers.GetAllGroups)
	r.Get("/groups/:id", controllers.GetGroupWithProducts)
	r.Delete("/groups/:id", controllers.DeleteGroup)
} 