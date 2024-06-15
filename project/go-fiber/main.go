package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "index.html",
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	app.Get("/api/", rootHandler)
	app.Get("/api/param/:parameter", parameterHandler)
	app.Get("/api/sub/:parameter?", optionalParameterHandler)

	app.Post("/api/create", postDataHandler)
	app.Listen("localhost:3000")
}

func rootHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func parameterHandler(c *fiber.Ctx) error {
	parameter := c.Params("parameter") // string
	return c.SendString("value of parameter = " + parameter)
}

func optionalParameterHandler(c *fiber.Ctx) error {
	if c.Params("parameter") != "" {
		return c.SendString("parameter is" + c.Params("name"))
	}
	return c.SendString("no parameter")
}
func postDataHandler(c *fiber.Ctx) error {
	payload := struct {
		ProjectName string `json:"projectName"`
		SoldToCode  string `json:"soldToCode"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	return c.SendString("")
}
