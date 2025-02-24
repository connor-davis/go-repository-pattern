package todos

import (
	"github.com/connor-davis/repository-pattern/internal/postgres"
	"github.com/gofiber/fiber/v2"
)

// CreateTodo godoc
// @Summary      Create a new todo
// @Description  Create a new todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo body      postgres.Todo true "Todo object"
// @Success      201  {object}  postgres.Todo
// @Failure      400  {string}  string "Invalid request payload"
// @Failure      500  {string}  string "Failed to create todo"
// @Router       /todos [post]
func (t *Todos) CreateTodo(c *fiber.Ctx) error {
	var todo postgres.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	todo, err := t.todosService.CreateTodo(c.Context(), todo.Title, todo.Description)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(todo)
}
