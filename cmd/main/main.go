package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", GetStartPage)

	log.Fatal(app.Listen(":3000"))

}

type MainPageData struct{
	Id uint `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

func GetStartPage(c *fiber.Ctx) error {

	data := MainPageData{
		Id: 12345,
		Title: "Hello World",
		Body: "Lorem dogh",
	}

	return c.JSON(data)
}