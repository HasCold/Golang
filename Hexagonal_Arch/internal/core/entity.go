// internal/core/entity.go
package core

type User struct {
	ID    int
	Name  string
	Email string
}

// SRP (Single Responsibility): User struct only represents data.
