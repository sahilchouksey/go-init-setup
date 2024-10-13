package main

import (
	"github.com/sahilchouksey/go-init-setup/app"
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
