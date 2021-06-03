package main

import (
	"log"

	ic "github.com/branogarbo/imgcli/util"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		GETOnly: true,
	})

	app.Get("/gen", func(c *fiber.Ctx) error {
		if len(c.Body()) == 0 {
			return c.Status(fiber.StatusBadRequest).SendString("provide bytes or a URL in the request body")
		}

		icConfig := struct {
			Mode   string
			Width  int
			Chars  string
			Invert bool
			Web    bool
		}{}

		err := c.QueryParser(&icConfig)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		img, err := ic.OutputImage(ic.OutputConfig{
			Src:          string(c.Body()),
			OutputMode:   icConfig.Mode,
			AsciiPattern: icConfig.Chars,
			OutputWidth:  icConfig.Width,
			IsInverted:   icConfig.Invert,
			IsUseWeb:     icConfig.Web,
			IsSrcBytes:   !icConfig.Web,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.SendString(img)
	})

	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	log.Fatal(app.Listen(":3000"))
}
