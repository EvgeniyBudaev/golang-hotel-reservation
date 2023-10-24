package api

import (
	"fmt"
	"github.com/EvgeniyBudaev/golang-hotel-reservation/internal/models"
	"github.com/gofiber/fiber/v2"
)

func getAuthUser(c *fiber.Ctx) (*models.User, error) {
	user, ok := c.Context().UserValue("user").(*models.User)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return user, nil
}
