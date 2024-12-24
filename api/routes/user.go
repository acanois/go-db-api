package routes

import "github.com/acanois/dbapi/internal/database"

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
}

func AddUser(u *User) {
	database.Connect()
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
