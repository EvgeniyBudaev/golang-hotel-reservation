package api

import (
	"github.com/EvgeniyBudaev/golang-hotel-reservation/internal/models"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := models.User{
		FirstName: "James",
		LastName:  "Bob",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("James")
}
