// internal/ports/repository.go
package ports

import "hexagonal-architecture/internal/core"

type UserRepository interface {
	Save(user *core.User) error
	FindByID(id int) (*core.User, error)
	FindAll() ([]*core.User, error)
}

// DIP (Dependency Inversion): Business logic depends on interfaces, not concrete implementations.
