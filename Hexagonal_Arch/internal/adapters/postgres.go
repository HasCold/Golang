// internal/adapters/postgres.go
package adapters

// Implement Adapters (Infrastructure Layer)

import (
	"database/sql"
	"errors"
	"hexagonal-architecture/internal/core"
	// _ "github.com/lib/pq"
)

type PostgresUserRepo struct {
	db *sql.DB
}

func NewPostgresUserRepo(db *sql.DB) *PostgresUserRepo {
	return &PostgresUserRepo{db: db}
}

func (r *PostgresUserRepo) Save(user *core.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, name, email) VALUES ($1, $2, $3)", user.ID, user.Name, user.Email)
	return err
}

func (r *PostgresUserRepo) FindByID(id int) (*core.User, error) {
	var user core.User
	err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// LSP (Liskov Substitution): The system works regardless of storage type (PostgreSQL or in-memory).
// DIP (Dependency Inversion): Business logic doesn’t change if we switch databases.
