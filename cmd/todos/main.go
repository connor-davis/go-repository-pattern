package main

import (
	"context"

	"github.com/connor-davis/repository-pattern/internal/api"
	"github.com/connor-davis/repository-pattern/internal/postgres"
	"github.com/connor-davis/repository-pattern/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:Connor100203@localhost:5432/todos-go?sslmode=disable")

	if err != nil {
		return err
	}

	defer conn.Close(ctx)

	// Initialize the postgres queries
	queries := postgres.New(conn)

	// Initialize the services
	todosService := services.NewTodosService(queries)

	// Initialize the Fiber app
	app := fiber.New(fiber.Config{
		AppName:           "Todos API",
		ReduceMemoryUsage: true,
		Prefork:           true,
	})

	app.Use(logger.New())

	// Initialize the API
	api := api.NewApi(app, *todosService)
	api.Setup()

	return app.Listen(":3000")
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
