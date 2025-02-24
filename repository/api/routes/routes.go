package routes

import (
	"github.com/connor-davis/repository-pattern/repository/api/routes/todos"
	"github.com/connor-davis/repository-pattern/repository/services"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	router       fiber.Router
	todosService services.TodosService
}

func NewRoutes(router fiber.Router, todosService services.TodosService) *Routes {
	return &Routes{
		router:       router,
		todosService: todosService,
	}
}

func (r *Routes) Setup() {
	todos := todos.NewTodos(r.router, r.todosService)
	todos.Setup()
}
