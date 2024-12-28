package routes

import (
	"fmt"

	"github.com/acanois/dbapi/internal/database"
)

type User struct {
	ID        int64
	UserName  string
	FirstName string
	LastName  string
	Email     string
}

func AddUser(u *User) (int64, error) {
	db, err := database.Connect()

	if err != nil {
		return 0, fmt.Errorf("AddUser: %v", err)
	}

	result, err := db.Exec(
		"INSERT INTO user (first_name, last_name, email) VALUES (?, ?, ?)",
		u.FirstName,
		u.LastName,
		u.Email,
	)

	if err != nil {
		return 0, fmt.Errorf("AddUser: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddUser: %v", err)
	}

	return id, nil
}

func UpdateUser(u *User) {
	database.Connect()
}

func GetUser(u *User) {
	database.Connect()
}

func DeleteUser(u *User) {
	database.Connect()
}
