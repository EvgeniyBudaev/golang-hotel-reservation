package middleware

import (
	"fmt"
	"github.com/EvgeniyBudaev/golang-hotel-reservation/internal/models"
	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*models.User)
	if !ok {
		return fmt.Errorf("not authorized")
	}
	if !user.IsAdmin {
		return fmt.Errorf("not authorized")
	}
	return c.Next()
}
