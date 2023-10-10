package main

import (
	"github.com/TVM-LLC/docapp_backend/app"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	// setup and run app
	err := app.SetupAndRunServer()
	if err != nil {
		log.Trace(err)
		panic(err)
	}
}
