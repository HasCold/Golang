package repository

import "database/sql"

type User struct {
	ID   uint
	Name string
}

func GetUserByID(db *sql.DB, id int) (User, error) {
	var user User
	row := db.QueryRow("SELECT id, name FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
