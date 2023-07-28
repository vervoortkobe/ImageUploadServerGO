package handlers

import (
	"fmt"
	"server/exports"

	"github.com/gofiber/fiber/v2"
)

func AuthHandler(c *fiber.Ctx) error {
	payload := exports.UserCreds{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	fmt.Printf(payload.Username + ": " + payload.Password + "\n")
	return c.JSON(payload)
}
