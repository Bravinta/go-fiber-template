package interfaces

import "github.com/gofiber/fiber/v2"

func TestInterface(c *fiber.Ctx) error {
	return c.SendString("Test")
}
