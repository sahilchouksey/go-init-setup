package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/sahilchouksey/go-init-setup/database"
	"github.com/sahilchouksey/go-init-setup/handlers"
	todo_handlers "github.com/sahilchouksey/go-init-setup/handlers/todo"
	"github.com/sahilchouksey/go-init-setup/utils"
)

func SetupRoutes(app *fiber.App, store *database.PostgreSQLStore) {
	// Health check endpoint
	app.Get("/ping", utils.MakeHTTPHandleFunc(handlers.HandleCheckHealth, store))

	// todo endpoints
	app.Get("/todos", utils.MakeHTTPHandleFunc(todo_handlers.GetAllTodos, store))
	//
	app.Get("/add/todo", utils.MakeHTTPHandleFunc(todo_handlers.AddTodoHandler, store))
	//
	app.Get("/update/todo", utils.MakeHTTPHandleFunc(todo_handlers.UpdateTodoHandler, store))
	//
	app.Get("/delete/todo", utils.MakeHTTPHandleFunc(todo_handlers.DeleteTodoHandler, store))
}
