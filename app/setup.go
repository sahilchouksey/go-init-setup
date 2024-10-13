package app

import (
	"os"

	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/sahilchouksey/go-init-setup/api"
	"github.com/sahilchouksey/go-init-setup/config"
	"github.com/sahilchouksey/go-init-setup/router"

	//"github.com/sahilchouksey/go-init-setup/databa/go-init-setup/router"
	"github.com/sahilchouksey/go-init-setup/database"
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
	app.Use(logger.New())

	app.Use(recover.New())

	// Setup Routes
	router.SetupRoutes(app, store)

	// Attach Swagger

	// Get the PORT & Start the Server
	return server.Run()

}
