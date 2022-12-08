package router

import (
	"github.com/gofiber/fiber/v2"
	v1 "go-fiber-template/internal/infrastructure/router/v1"
)

func Routes(r *fiber.App) {
	versionOne := r.Group("/v1")
	{
		v1.Routes(versionOne)
	}
}
