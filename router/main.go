package router

import (
	"github.com/TVM-LLC/docapp_backend/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	app.GET("/ping", handlers.HandleCheckHealth)
}
