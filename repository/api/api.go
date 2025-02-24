package api

import (
	"github.com/connor-davis/repository-pattern/repository/api/routes"
	"github.com/connor-davis/repository-pattern/repository/services"
	"github.com/gofiber/fiber/v2"
)

type Api struct {
	router       fiber.Router
	todosService services.TodosService
}

func NewApi(router fiber.Router, todosService services.TodosService) *Api {
	return &Api{
		router:       router,
		todosService: todosService,
	}
}

func (a *Api) Setup() {
	r := routes.NewRoutes(a.router, a.todosService)
	r.Setup()
}
