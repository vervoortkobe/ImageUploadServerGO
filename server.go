package main

import (
	"fmt"
	"log"
	"server/dbactions"
	"server/exports"
	"server/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	PORT := ":3000"

	exports.MongoClient()
	fmt.Print("💽 | Connected to MongoDB!\n")

	dbactions.LogAll()

	exports.App.Static("/", "./public")

	exports.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	exports.App.Get("/all", handlers.GetJsonHandler)

	exports.App.Post("/post", func(c *fiber.Ctx) error {
		payload := struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}
		fmt.Printf(payload.Name + ": " + payload.Email + "\n")
		return c.JSON(payload)
	})

	exports.App.Post("/upload", handlers.UploadHandler)

	exports.App.Get("/:value", handlers.ImageHostBuilder)

	exports.App.Get("*", func(c *fiber.Ctx) error {
		return c.Redirect("/")
	})

	///////////////////////////////////////////////////////////////////////

	fmt.Printf("⚡ | WebServer listening on [http://localhost%s]!\n", PORT)
	log.Fatal(exports.App.Listen(PORT))
}
