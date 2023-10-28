package model

import (
	"fmt"
	"user-service/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func GetUserById(id int) (*User, *error) {
	var user User
	db := utils.GetConnector()

	if result := db.First(&user, id); result.Error != nil {
		return nil, &result.Error
	}
	return &user, nil
}

func GetUserByFirstName(firstName string) (*User, error) {
	var user User
	db := utils.GetConnector()

	if result := db.First(&user, "firstName = ?", firstName); result.Error != nil {
		err := NotFoundError{
			ErrorMessage: result.Error.Error(),
			Table:        "User",
			Val:          fmt.Sprintf("firstName = %s", firstName),
		}
		return nil, &err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	db := utils.GetConnector()

	if result := db.First(&user, "email = ?", email); result.Error != nil {
		err := NotFoundError{
			ErrorMessage: result.Error.Error(),
			Table:        "User",
			Val:          fmt.Sprintf("email = %s", email),
		}
		return nil, &err
	}
	return &user, nil
}

func CreateUser(user *User) error {
	db := utils.GetConnector()

	if result := db.Create(user); result.RowsAffected >= 1 {
		return result.Error
	}
	return nil
}
