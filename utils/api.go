package utils

import (
	"github.com/sahilchouksey/go-init-setup/database"
	"github.com/gin-gonic/gin"
)

func MakeHTTPHandleFunc(handler func(c *gin.Context, store *database.PostgreSQLStore) error, store *database.PostgreSQLStore) func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := handler(c, store); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}
}
func WriteJSON(c *gin.Context, status int, data interface{}, headers map[string]string, err error) {
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	if len(headers) > 0 {
		for key, value := range headers {
			c.Header(key, value)
		}
	} else {
		c.Header("Content-Type", "application/json")
	}

	c.JSON(status, data)

}
