package main

import (
	"log"

	ic "github.com/branogarbo/imgcli/util"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		UnescapePath: true,
	})

	app.All("/gen", func(c *fiber.Ctx) error {
		if len(c.Body()) == 0 {
			return c.Status(fiber.StatusBadRequest).SendString("provide a URL in the request body")
		}

		icConfig := struct {
			Mode   string
			Width  int
			Chars  string
			Invert bool
		}{
			Mode:  "ascii",
			Width: 100,
			Chars: " .:-=+*#%@",
		}

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
			IsUseWeb:     true, // only web for now
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.SendString(img)
	})

	app.All("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	log.Fatal(app.Listen(":3000"))
}
