package handlers

import "github.com/gin-gonic/gin"

func HandleCheckHealth(c *gin.Context) {
	c.String(200, "PING")
}
