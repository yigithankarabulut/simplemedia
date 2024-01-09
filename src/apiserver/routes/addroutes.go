package routes

import "github.com/gofiber/fiber/v2"

type Handler interface {
	AddRoutes(router fiber.Router)
}
