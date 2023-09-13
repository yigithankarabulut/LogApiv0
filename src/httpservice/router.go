package httpservice

import "github.com/gofiber/fiber/v2"

func Router(r fiber.Router) {
	r.Post("/admin", Register)
	r.Post("/login", Login)
	r.Post("/log", Log)
	r.Delete("/admin", DeleteUser)
	r.Put("/register", UpdateUser)
}
