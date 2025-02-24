package todos

import (
	"github.com/connor-davis/repository-pattern/repository/services"
	"github.com/gofiber/fiber/v2"
)

type Todos struct {
	router       fiber.Router
	todosService services.TodosService
}

func NewTodos(router fiber.Router, todosService services.TodosService) *Todos {
	return &Todos{
		router:       router,
		todosService: todosService,
	}
}

func (t *Todos) Setup() {
	t.router.Get("/todos", t.GetTodos)
	t.router.Get("/todos/:id", t.GetTodoById)
	t.router.Post("/todos", t.CreateTodo)
	t.router.Put("/todos/:id", t.UpdateTodoById)
	t.router.Delete("/todos/:id", t.DeleteTodoById)
}
