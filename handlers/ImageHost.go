package handlers

import (
	"fmt"
	"server/dbactions"
	"server/exports"

	"github.com/gofiber/fiber/v2"
)

func ImageHost(value string) (exports.Image, error) {
	img, err := dbactions.FindImage(value)
	if img != exports.EmptyImage && err == nil {
		return img, nil
	}
	return exports.EmptyImage, nil
}

func ImageHostInefficient(c *fiber.Ctx) error {
	results, err := dbactions.GetAll()
	if err != nil {
		return c.SendString(fmt.Sprint(err))
	}

	for i, r := range results {
		_ = i
		if c.Params("value") == r.Id {
			c.Set("Content-Type", "image/jpeg")
			return c.SendString(r.Id)
		}
	}
	return c.Redirect("/")
}
