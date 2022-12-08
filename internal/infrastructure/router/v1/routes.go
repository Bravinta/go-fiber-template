package v1

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-template/internal/interfaces"
)

func Routes(v1 fiber.Router) {
	v1.Get("/test", interfaces.TestInterface)
}
