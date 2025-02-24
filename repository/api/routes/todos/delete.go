package todos

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// DeleteTodoById godoc
// @Summary      Delete Todo by ID
// @Description  Delete a todo by ID
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      int     true  "Todo ID"
// @Success      200  {string}  string "ok"
// @Failure      400  {string}  string "Invalid ID"
// @Failure      404  {string}  string "Todo not found"
// @Failure      500  {string}  string "Internal server error"
// @Router       /todos/{id} [delete]
func (t *Todos) DeleteTodoById(c *fiber.Ctx) error {
	id := c.Params("id")

	todoID, err := strconv.Atoi(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	err = t.todosService.DeleteTodoById(c.Context(), int32(todoID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("ok")
}
