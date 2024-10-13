package router

import (
	"github.com/sahilchouksey/go-init-setup/database"
	"github.com/sahilchouksey/go-init-setup/handlers"
	todo_handlers "github.com/sahilchouksey/go-init-setup/handlers/todo"
	utils "github.com/sahilchouksey/go-init-setup/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine, store *database.PostgreSQLStore) {
	// Health check endpoint
	app.GET("/ping", utils.MakeHTTPHandleFunc(handlers.HandleCheckHealth, store))

	// todo endpoints
	app.GET("/todos", utils.MakeHTTPHandleFunc(todo_handlers.GetAllTodos, store))

	app.GET("/add/todo", utils.MakeHTTPHandleFunc(todo_handlers.AddTodoHandler, store))

	app.GET("/update/todo", utils.MakeHTTPHandleFunc(todo_handlers.UpdateTodoHandler, store))

	app.GET("/delete/todo", utils.MakeHTTPHandleFunc(todo_handlers.DeleteTodoHandler, store))
}
