// internal/adapters/inmem.go
package adapters

// Implement Adapters (Infrastructure Layer)

import (
	"errors"
	"hexagonal-architecture/internal/core"
	"sync"
)

type InMemoryUserRepo struct {
	users map[int]*core.User
	mu    sync.Mutex
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[int]*core.User),
	}
}

func (r *InMemoryUserRepo) Save(user *core.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepo) FindByID(id int) (*core.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepo) FindAll() ([]*core.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	users := []*core.User{}
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}

// LSP (Liskov Substitution): The system works regardless of storage type (PostgreSQL or in-memory).
// DIP (Dependency Inversion): Business logic doesn’t change if we switch databases.
