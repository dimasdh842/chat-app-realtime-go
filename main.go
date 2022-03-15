package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   "1310843",
		Key:     "1886f5a1c98a16d70491",
		Secret:  "47c26938e75711c566a8",
		Cluster: "ap1",
		Secure:  true,
	}

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		pusherClient.Trigger("chat", "message", data)

		return c.JSON([]string{})
	})

	app.Get("test", func(c *fiber.Ctx) error {
		return c.JSON("hello")
	})

	app.Listen(":8000")
}
