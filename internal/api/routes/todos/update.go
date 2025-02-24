package todos

import (
	"strconv"

	"github.com/connor-davis/repository-pattern/internal/postgres"
	"github.com/gofiber/fiber/v2"
)

// UpdateTodo godoc
// @Summary      Update a todo
// @Description  Update a todo by ID
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      int     true  "Todo ID"
// @Param        todo body      postgres.Todo true "Todo object"
// @Success      200  {object}  postgres.Todo
// @Failure      400  {string}  string "Invalid ID"
// @Failure      404  {string}  string "Todo not found"
// @Failure      500  {string}  string "Internal server error"
// @Router       /todos/{id} [put]
func (t *Todos) UpdateTodoById(c *fiber.Ctx) error {
	id := c.Params("id")

	todoID, err := strconv.Atoi(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	var todo postgres.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	if err := t.todosService.UpdateTodoById(c.Context(), int32(todoID), todo.Title, todo.Description); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("ok")
}
