// main.go
package main

import (
	"fmt"
	"hexagonal-architecture/internal/adapters"
	"hexagonal-architecture/internal/core"
)

func main() {
	// Choose the repository (In-memory or Postgres)
	repo := adapters.NewInMemoryUserRepo() // Switch to PostgreSQL if needed

	// Create Service
	service := core.NewUserService(repo)

	// Create and Fetch a User
	user := &core.User{ID: 1, Name: "John Doe", Email: "john@example.com"}
	service.CreateUser(user)

	retrievedUser, _ := service.GetUser(1)
	fmt.Println("Retrieved User:", retrievedUser.Name)
}

// OCP (Open-Closed Principle): Can switch repositories easily.
