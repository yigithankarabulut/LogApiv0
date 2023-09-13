package httpservice

import (
	"LogApi00/src/handlers"
	"LogApi00/src/storage"
	"github.com/gofiber/fiber/v2"
)

func DeleteUser(c *fiber.Ctx) error {
	data := make(map[string]string)
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("body not parse")
	}
	if handlers.JwtAccess(data["jwt"], true) == false {
		return c.Status(fiber.StatusUnauthorized).SendString("invalid jwt")
	}
	if err := storage.DeleteUser(data["username"]); err != nil {
		return c.Status(fiber.StatusNotFound).SendString("user not found")
	}
	return c.Status(200).SendString("user delete successfully!")
}
