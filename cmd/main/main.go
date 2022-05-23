package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/strikermcs/dresscode/pkg/config"
	"github.com/strikermcs/dresscode/pkg/controllers"
	"github.com/strikermcs/dresscode/pkg/routes"
)

func main() {
	app := fiber.New()
	db := config.GetDb()
	controllers.InitDatabase(db)
	routes.InitRoutes(app)
	log.Fatal(app.Listen(":3000"))
}