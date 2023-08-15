package handler

import (
	"log"

	"alfredtomo.dev/todos-app/database"
	"alfredtomo.dev/todos-app/model/entity"
	"github.com/gofiber/fiber/v2"
)

func GetTodoHandler(ctx *fiber.Ctx) error {
	var todo []entity.Todo

	result := database.DB.Find(&todo)

	if result.Error != nil {
		log.Println(result.Error)
	}

	// Other way to extract error
	// err := database.DB.Find(&todo).Error
	// if err != nil {
	// 	log.Println(err)
	// }

	return ctx.JSON(todo)
}

func CreateTodoHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"todo": "here we are",
	})
}

func UpdateTodoHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"todo": "here we are",
	})
}

func CompleteTodoHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"todo": "here we are",
	})
}

func DeleteTodoHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"todo": "here we are",
	})
}
