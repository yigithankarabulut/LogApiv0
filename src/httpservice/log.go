package httpservice

import (
	"LogApi00/src/handlers"
	"context"
	"github.com/gofiber/fiber/v2"
	"time"
)

func Log(c *fiber.Ctx) error {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data := make(map[string]string)
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Body Not Parse")
	}
	if handlers.JwtAccess(data["jwt"], false) == false {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid JWT")
	}
	username, jwtErr := handlers.JwtGetUsername(data["jwt"])
	if jwtErr != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(jwtErr.Error())
	}
	err := handlers.CreateDataFile(data["data"], username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("log files not created")
	}

	return c.Status(200).SendString("log files create successfully!")
}
