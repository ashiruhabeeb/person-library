package handler

import "github.com/gofiber/fiber/v2"

func HealthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"status": "Server is healthy"})
}