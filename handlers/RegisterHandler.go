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

	username, err := dbactions.Register(payload.Username, payload.Password)
	if username == "err_duplicate_username" && err != nil {
		return c.SendString("err_duplicate_username")
	} else if username == "err_insert_one" && err != nil {
		return c.SendString("err_insert_one")
	} else {
		return c.SendString(fmt.Sprint(username))
	}
}
