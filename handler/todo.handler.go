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

// GetTodo godoc
// @Summary      Get All Todo's
// @Description  get string by ID
// @Tags         todo
// @Produce      json
// @Success      200  {object}  []entity.Todo
// @Failure      400  {object}  fiber.Error
// @Router       /todo [get]
func GetTodoHandler(ctx *fiber.Ctx) error {
	var todo []entity.Todo

	result := database.DB.Find(&todo)

	if result.Error != nil {
		log.Println(result.Error)
		return fiber.NewError(400, "failed")
	}

	// Other way to extract error
	// err := database.DB.Find(&todo).Error
	// if err != nil {
	// 	log.Println(err)
	// }

	return ctx.JSON(todo)
}

// CreateTodo godoc
// @Summary      Create a new Todo
// @Description  Create a new Todo with the input payload
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param        request body request.CreateTodoRequest true "Payload for creating a new todo"
// @Success      200  {object}  entity.Todo
// @Failure      400  {object}  fiber.Error
// @Router       /todo [post]
func CreateTodoHandler(ctx *fiber.Ctx) error {
	todo := new(request.CreateTodoRequest)

	if err := ctx.BodyParser(todo); err != nil {
		return err
	}

	if errValidate := validate.Struct(todo); errValidate != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "failed")
	}

	newTodo := entity.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   false,
	}

	result := database.DB.Create(&newTodo)

	if result.Error != nil {
		log.Println(result.Error)
		return fiber.NewError(fiber.ErrBadRequest.Code, "failed")
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newTodo,
	})
}

// UpdateTodo godoc
// @Summary      Update a Todo
// @Description  Update a Todo with params id and input payload
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param        id path string true "Todo ID"
// @Param        request body request.UpdateTodoRequest true "Payload for creating a new todo"
// @Success      200  {object}  entity.Todo
// @Failure      400  {object}  fiber.Error
// @Router       /todo/:id [patch]
func UpdateTodoHandler(ctx *fiber.Ctx) error {
	todoId := ctx.Params("id")

	todo := new(request.UpdateTodoRequest)

	if err := ctx.BodyParser(todo); err != nil {
		return err
	}

	if errValidate := validate.Struct(todo); errValidate != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, errValidate.Error())
	}

	var todoEntity entity.Todo

	result := database.DB.First(&todoEntity, todoId)

	if result.Error != nil {
		log.Println(result.Error)
		return fiber.NewError(fiber.ErrBadRequest.Code, "failed")
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
		return fiber.NewError(fiber.ErrBadRequest.Code, "failed")
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    todoEntity,
	})
}

// CompleteTodo godoc
// @Summary      Complete a Todo
// @Description  Complete a Todo with params id
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param        id path string true "Todo ID"
// @Success      200  {object}  entity.Todo
// @Failure      400  {object}  fiber.Error
// @Router       /todo/:id/complete [post]
func CompleteTodoHandler(ctx *fiber.Ctx) error {
	todoId := ctx.Params("id")

	var todoEntity entity.Todo

	result := database.DB.First(&todoEntity, todoId)

	if result.Error != nil {
		log.Println(result.Error)
		return fiber.NewError(fiber.ErrBadRequest.Code, "failed")
	}

	todoEntity.Completed = true

	dbError := database.DB.Save(&todoEntity).Error

	if dbError != nil {
		log.Println(dbError)
		return fiber.NewError(fiber.ErrBadRequest.Code, "failed")
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    todoEntity,
	})
}

// DeleteTodo godoc
// @Summary      Delete a Todo
// @Description  Delete a Todo with params id
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param        id path string true "Todo ID"
// @Success      200  {object}  string
// @Failure      400  {object}  fiber.Error
// @Router       /todo/:id [delete]
func DeleteTodoHandler(ctx *fiber.Ctx) error {
	todoId := ctx.Params("id")

	var todo entity.Todo

	result := database.DB.First(&todo, todoId)

	if result.Error != nil {
		log.Println(result.Error)
		return fiber.NewError(fiber.ErrBadRequest.Code, "failed")
	}

	dbError := database.DB.Delete(&todo).Error

	if dbError != nil {
		log.Println(dbError)
		return fiber.NewError(fiber.ErrBadRequest.Code, "failed")
	}

	return ctx.Status(200).SendString("success")
}
