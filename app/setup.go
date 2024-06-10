package app

import (
	"fmt"
	"os"
	"time"

	"github.com/TVM-LLC/docapp_backend/api"
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
	store, err := database.Start()
	if err != nil {
		print("Check whether the Postgres is running or not\n")
		print("If not running, run the following command\n")
		print("bash ./docker_essentials\n")
		return err
	}

	if err := store.Init(); err != nil {
		print("Check whether the Postgres is running or not\n")
		print("If not running, run the following command\n")
		print("bash ./docker_essentials\n")
		return err
	}

	// Defer Closing DB
	defer store.Close()

	// Init API
	var server *api.APIServer = api.NewAPIServer(os.Getenv("PORT"))
	app := server.GetEngine()

	// Attach Middleware
	// Custom Logger
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// custom format
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

	app.Use(gin.Recovery())

	// Setup Routes
	router.SetupRoutes(app, store)

	// Attach Swagger

	// Get the PORT & Start the Server
	return server.Run()

}
