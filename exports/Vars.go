package exports

import "github.com/gofiber/fiber/v2"

var App *fiber.App = fiber.New()

type Image struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Data      string `json:"data"`
	Timestamp int    `json:"timestamp"`
}

var EmptyImage Image = Image{}

type UserCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
