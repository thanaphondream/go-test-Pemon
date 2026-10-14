package http

import (
	"github.com/gofiber/fiber/v2"
	"test/service"
	"os"
	// "log"
)

type httpViewMethop interface {
	ViewGet(c *fiber.Ctx) error
	ViewGet_DC_1_Html(c *fiber.Ctx) error
	ViewGet_AC_2_Html(c *fiber.Ctx) error
}

type httpViewServices struct {
	service service.ServiceView
}

func NewHttpView(s service.ServiceView) httpViewMethop{
	return httpViewServices{service: s}
}

func (s httpViewServices) ViewGet(c *fiber.Ctx) error {
	view, err := s.service.Views()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(view)
}

func (s httpViewServices) ViewGet_DC_1_Html(c *fiber.Ctx) error {
	urlh := string("./templates/DC_1.html")
	data, err := os.ReadFile(urlh)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("ไม่สามารถอ่านไฟล์ได้")
	}

	htmlContent := string(data)

	// log.Println(htmlContent) 

	return c.Type("html").SendString(htmlContent)
}

func (s httpViewServices) ViewGet_AC_2_Html(c *fiber.Ctx) error {
	urlh := string("./templates/AC_2.html")
	data, err := os.ReadFile(urlh)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("ไม่สามารถอ่านไฟล์ได้")
	}

	htmlContent := string(data)

	// log.Println(htmlContent) 

	return c.Type("html").SendString(htmlContent)
}

