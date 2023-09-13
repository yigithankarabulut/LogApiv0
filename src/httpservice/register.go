package httpservice

import (
	"LogApi00/src/handlers"
	"LogApi00/src/models"
	"LogApi00/src/storage"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	data := make(map[string]string)
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if handlers.JwtAccess(data["jwt"], true) == false {
		return c.Status(fiber.StatusBadGateway).SendString("Invalid JWT")
	}
	userData := models.User{Username: data["username"], Password: data["password"]}

	if err := storage.CreateUser(userData); err != nil {
		return c.Status(fiber.StatusBadGateway).SendString("Create Failed!")
	}
	return c.Status(200).SendString("Create Successfully!")
}
