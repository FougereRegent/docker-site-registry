package dto

import (
	"user-service/model"
)

type UserConnection struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInformation struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserCreate struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func UserToUserInformation(user model.User) UserInformation {
	return UserInformation{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email}
}

func UserConnectionToUser(user UserConnection) model.User {
	return model.User{Email: user.Email, Password: user.Password}
}

func UserInformationToUser(user UserInformation) model.User {
	return model.User{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email}
}
