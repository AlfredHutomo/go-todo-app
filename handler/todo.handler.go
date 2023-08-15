package handler

import (
	"log"

	"alfredtomo.dev/todos-app/database"
	"alfredtomo.dev/todos-app/model/entity"
	"alfredtomo.dev/todos-app/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

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
	todo := new(request.CreateTodoRequest)

	if err := ctx.BodyParser(todo); err != nil {
		return err
	}

	if errValidate := validate.Struct(todo); errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": errValidate.Error(),
		})
	}

	newTodo := entity.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		Completed: false,
	}

	result := database.DB.Create(&newTodo)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": newTodo,
	})
}

func UpdateTodoHandler(ctx *fiber.Ctx) error {
	todoId := ctx.Params("id")

	todo := new(request.UpdateTodoRequest)

	if err := ctx.BodyParser(todo); err != nil {
		return err
	}

	if errValidate := validate.Struct(todo); errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": errValidate.Error(),
		})
	}

	var todoEntity entity.Todo

	result := database.DB.First(&todoEntity, todoId)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
		})
	}

	if todo.Title != "" {
		todo.Title = todoEntity.Title
	}

	if todo.Description != "" {
		todo.Description = todoEntity.Description
	}

	dbError := database.DB.Save(&todoEntity).Error

	if dbError != nil {
		log.Println(dbError)
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": todoEntity,
	})
}

func CompleteTodoHandler(ctx *fiber.Ctx) error {
	todoId := ctx.Params("id")

	var todoEntity entity.Todo

	result := database.DB.First(&todoEntity, todoId)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
		})
	}

	todoEntity.Completed = true

	dbError := database.DB.Save(&todoEntity).Error

	if dbError != nil {
		log.Println(dbError)
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data": todoEntity,
	})
}

func DeleteTodoHandler(ctx *fiber.Ctx) error {
	todoId := ctx.Params("id")

	var todo entity.Todo

	result := database.DB.First(&todo, todoId)

	if result.Error != nil {
		log.Println(result.Error)
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
		})
	}

	dbError := database.DB.Delete(&todo).Error

	if dbError != nil {
		log.Println(dbError)
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}
