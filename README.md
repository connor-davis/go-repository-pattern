# 🎉 Repository Pattern in Go 🎉

## Project Overview

Welcome to the **Repository Pattern in Go** project! 🚀 This project showcases the implementation of the Repository Pattern in Go. The Repository Pattern is a design pattern that mediates data from and to the Domain and Data Mapping layers. Repositories encapsulate the logic required to access data sources, centralizing common data access functionality and providing better maintainability. This decouples the data access logic from the business logic, making your code cleaner and more modular.

## Installation

Ready to dive in? Follow these steps to get started:

1. **Clone the repository** to your local machine:

    ```sh
    git clone https://github.com/connor-davis/repository-pattern.git
    ```

2. **Navigate** to the project directory:

    ```sh
    cd repository-pattern
    ```

3. **Install** the necessary dependencies:

    ```sh
    go mod tidy
    ```

## Usage

Let's see the Repository Pattern in action! You can run the examples by executing the following command:

```sh
go run cmd/todos/main.go
```

### Project Structure

Here's a quick tour of the project structure:

- `cmd/todos/main.go`: The entry point of the application.
- `repository/`: Contains the repository interfaces and their implementations.
  - `postgres/`: Contains the PostgreSQL specific implementations.
    - `db.go`: Contains the database connection and query methods.
- `services/`: Contains the business logic that uses the repositories.
  - `todos.go`: Contains the implementation of the todos service.

### Example

Let's walk through a simple example of how to use the repository pattern in this project:

```go
package main

import (
	"context"
	"fmt"
	"github.com/connor-davis/repository-pattern/repository/postgres"
	"github.com/connor-davis/repository-pattern/services"
)

func main() {
	// Initialize the postgres queries
	queries := postgres.New(nil) // Replace nil with actual DB connection

	// Create a new service
	todosService := services.NewTodosService(queries)

	// Perform operations
	ctx := context.Background()
	todosService.CreateTodo(ctx, "Sample Todo", "This is a sample todo")

	// Fetch and print todos
	todos, _ := todosService.GetTodos(ctx)
	fmt.Println(todos)
}
```

## Contributing

We love contributions! ❤️ If you'd like to contribute, please fork the repository and create a pull request with your changes. Make sure to follow the existing code style and include tests for any new functionality.

1. **Fork** the repository
2. **Create** a new branch (`git checkout -b feature-branch`)
3. **Commit** your changes (`git commit -am 'Add new feature'`)
4. **Push** to the branch (`git push origin feature-branch`)
5. **Create** a new Pull Request

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

Happy coding! 🎉
