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

	app.Get("/gen", func(c *fiber.Ctx) error {
		icConfig := struct {
			Src    string
			Mode   string
			Width  int
			Chars  string
			Invert bool
		}{
			Src:   "cdn.britannica.com/60/8160-050-08CCEABC/German-shepherd.jpg",
			Mode:  "ascii",
			Width: 100,
			Chars: " .:-=+*#%@",
		}

		err := c.QueryParser(&icConfig)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}

		img, err := ic.OutputImage(ic.OutputConfig{
			Src:          "http://" + icConfig.Src[1:len(icConfig.Src)-1],
			OutputMode:   icConfig.Mode,
			AsciiPattern: icConfig.Chars,
			OutputWidth:  icConfig.Width,
			IsInverted:   icConfig.Invert,
			IsUseWeb:     true, // only web for now
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.SendString(img)
	})

	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	log.Fatal(app.Listen(":3000"))
}
