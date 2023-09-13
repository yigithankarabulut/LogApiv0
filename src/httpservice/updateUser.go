package httpservice

import (
	"LogApi00/src/handlers"
	"LogApi00/src/storage"
	"github.com/gofiber/fiber/v2"
)

func UpdateUser(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("body not parse")
	}
	if handlers.JwtAccess(data["jwt"], false) == false {
		return c.Status(fiber.StatusUnauthorized).SendString("invalid jwt")
	}
	username, err := handlers.JwtGetUsername(data["jwt"])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}
	err2 := storage.UpdateUser(username, data["oldpwd"], data["newpwd"])
	if err2 != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err2.Error(),
		})
	}
	return c.Status(200).SendString("user update successfully!")
}
