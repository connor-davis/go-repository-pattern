package todos

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// GetTodos godoc
// @Summary     Get Todos
// @Description Get all todos
// @Tags        todos
// @Accept      json
// @Produce     json
// @Success     200 {array} postgres.Todo
// @Failure     404 {string} string "No todos found"
// @Failure     500 {string} string "Internal server error"
// @Router      /todos [get]
func (t *Todos) GetTodos(c *fiber.Ctx) error {
	todos, err := t.todosService.GetTodos(c.Context())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if todos == nil {
		return c.Status(fiber.StatusNotFound).SendString("No todos found")
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}

// GetTodoById godoc
// @Summary     Get Todo by ID
// @Description Get a todo by ID
// @Tags        todos
// @Accept      json
// @Produce     json
// @Param       id path int true "Todo ID"
// @Success     200 {object} postgres.Todo
// @Failure     400 {string} string "Invalid ID"
// @Failure     404 {string} string "Todo not found"
// @Failure	    500 {string} string "Internal server error"
// @Router      /todos/{id} [get]
func (t *Todos) GetTodoById(c *fiber.Ctx) error {
	id := c.Params("id")

	todoID, err := strconv.Atoi(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	todo, err := t.todosService.GetTodoById(c.Context(), int32(todoID))

	if strings.Contains(err.Error(), "no rows in result set") {
		return c.Status(fiber.StatusNotFound).SendString("Todo not found")
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}
