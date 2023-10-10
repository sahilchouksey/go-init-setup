package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/TVM-LLC/docapp_backend/config"
	"github.com/TVM-LLC/docapp_backend/database"
	"github.com/TVM-LLC/docapp_backend/router"
	"github.com/gin-gonic/gin"
)

func SetupAndRunServer() error {

	// Load ENV
	if err := config.LoadENV(); err != nil {
		return err
	}

	// Start DB
	db, err := database.Start()
	if err != nil {
		return err
	}
	if err := db.Init(); err != nil {
		return err
	}

	// Defer Closing DB
	db.Close()

	// Create app
	app := gin.New()

	// Attach Middleware
	// Custom Logger
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	//	r.Use(gin.Logger())
	app.Use(gin.Recovery())

	// Setup Routes
	router.SetupRoutes(app)

	// Attach Swagger

	// Get the PORT & Start the Server
	PORT := os.Getenv("PORT")
	// Listen and serve on 0.0.0.0:8080
	log.Println("Starting Server on PORT: " + PORT)
	app.Run(":" + PORT)

	return nil
}
