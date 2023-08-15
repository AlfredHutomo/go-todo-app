package routes

import (
	"alfredtomo.dev/todos-app/handler"
	"github.com/gofiber/fiber/v2"
)

func TodoRoutes(app *fiber.App){
	todo := app.Group("/todo")

	todo.Get("/", handler.GetTodoHandler)
	todo.Post("/", handler.CreateTodoHandler)
	todo.Put("/:id", handler.UpdateTodoHandler)
	todo.Delete("/:id", handler.DeleteTodoHandler)
}