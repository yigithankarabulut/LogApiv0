package httpservice

import (
	"LogApi00/src/handlers"
	"LogApi00/src/models"
	"LogApi00/src/storage"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Login(c *fiber.Ctx) error {
	var data models.User
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := storage.GetUser(data); err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("user not found")
	}
	err, tokenJwt := handlers.JwtCreate(data.Username, c)
	if err != nil {
		log.Fatal(err)
	}
	return c.Status(200).JSON(fiber.Map{
		"jwt": tokenJwt,
	})
}
