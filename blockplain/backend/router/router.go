package router

import (
    "github.com/gofiber/fiber/v2"
)

func SetupAndListen() {
    app := fiber.New()

    // Sample endpoint
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Blockplain Backend Running")
    })

    app.Listen(":3000")
}
