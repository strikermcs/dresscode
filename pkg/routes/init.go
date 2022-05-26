package routes

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(r *fiber.App){
	ProductRoutesInit(r)
	GroupsRouterInit(r)
}