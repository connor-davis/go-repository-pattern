package services

import (
	"context"

	"github.com/connor-davis/repository-pattern/repository/postgres"
)

type TodosService struct {
	queries *postgres.Queries
}

func NewTodosService(queries *postgres.Queries) *TodosService {
	return &TodosService{
		queries: queries,
	}
}

func (s *TodosService) GetTodos(ctx context.Context) ([]postgres.Todo, error) {
	return s.queries.GetTodos(ctx)
}

func (s *TodosService) GetTodoById(ctx context.Context, id int32) (postgres.Todo, error) {
	return s.queries.GetTodoById(ctx, id)
}

func (s *TodosService) CreateTodo(ctx context.Context, title, description string) (postgres.Todo, error) {
	return s.queries.CreateTodo(ctx, postgres.CreateTodoParams{
		Title:       title,
		Description: description,
	})
}

func (s *TodosService) UpdateTodoById(ctx context.Context, id int32, title, description string) error {
	return s.queries.UpdateTodoById(ctx, postgres.UpdateTodoByIdParams{
		ID:          id,
		Title:       title,
		Description: description,
	})
}

func (s *TodosService) DeleteTodoById(ctx context.Context, id int32) error {
	return s.queries.DeleteTodoById(ctx, id)
}
