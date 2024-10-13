package handlers

import (
	"github.com/sahilchouksey/go-init-setup/database"
	"github.com/gin-gonic/gin"
)

func HandleCheckHealth(c *gin.Context, store *database.PostgreSQLStore) error {
	c.String(200, "PONG")
	return nil
}
