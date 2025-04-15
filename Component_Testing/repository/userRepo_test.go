package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetUserByID(t *testing.T) {
	// We can mock the sql DB without running the actual sql query by the inbuilt library
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, name FROM users WHERE id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string("id", "name")).AddRows(1, "Alice"))

	userResp, err := GetUserByID(db, 1)
	if err != nil {
		t.Fatal(err)
	}

	if userResp.Name != "Alice" {
		t.Errorf("expected name Alice but got %s", userResp.Name)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("the requirements of expectation are not fulfilled %v", err)
	}
}
