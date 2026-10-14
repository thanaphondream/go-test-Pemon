package main

import (
	"fmt"
	"test/http"
	"test/repository"
	"test/service"

	"github.com/gofiber/fiber/v2"

)

func main() {
	fmt.Println("hello world")

	app := fiber.New()

	viewRepo := repository.NewRepoView()
	viewService := service.NewServiceView(viewRepo)
	viewAdanter := http.NewHttpView(viewService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/View1.js", viewAdanter.ViewGet)

	app.Get("/s/DC", viewAdanter.ViewGet_DC_1_Html)
	app.Get("/s/AC", viewAdanter.ViewGet_AC_2_Html)

	app.Listen(":3000")
}