package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
)


func main() {
    app := fiber.New()

	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID: "1663178",
		Key: "cb73b551b8495da100d8",
		Secret: "4c0e1e2ece696dc0912c",
		Cluster: "ap1",
		Secure: true,
	}


    app.Post("/api/message", func(c *fiber.Ctx) error {
		var data map[string]string
		err := c.BodyParser(&data)
		if err != nil {
			fmt.Println(err.Error())
		}
		pusherClient.Trigger("chat", "message", data)
        return c.SendString("Success")
    })

    app.Listen(":9090")
}