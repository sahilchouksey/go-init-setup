package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/sahilchouksey/go-init-setup/database"
)

func HandleCheckHealth(c fiber.Ctx, store *database.PostgreSQLStore) error {
	c.JSON(fiber.Map{"status": "ok"})
	return nil
}
