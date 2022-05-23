package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/strikermcs/dresscode/pkg/config"
	"github.com/strikermcs/dresscode/pkg/routes"
)

func main() {
	config.ConnectDb()
	app := fiber.New()
	routes.InitRoutes(app)
	log.Fatal(app.Listen(":3000"))
}