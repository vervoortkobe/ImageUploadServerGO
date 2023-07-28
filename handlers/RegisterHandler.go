package handlers

import (
	"fmt"
	"server/dbactions"
	"server/exports"

	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(c *fiber.Ctx) error {

	payload := exports.UserCreds{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	fmt.Printf("New user registered > " + payload.Username + ": " + payload.Password + "\n")

	username, err := dbactions.Register(payload.Username, payload.Password)
	if username == "err_duplicate_username" && err != nil {
		return nil //err
	} else if username == "err_insert_one" && err != nil {
		return nil //err
	} else {
		return c.Redirect(fmt.Sprintf("/%s", username))
	}
}
