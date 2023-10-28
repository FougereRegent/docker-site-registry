package dto

import (
	"user-service/model"
)

type UserConnection struct {
	Email    string
	Password string
}

type UserInformation struct {
	FirstName string
	LastName  string
	Email     string
}

func UserToUserConnection(user model.User) UserConnection {
	return UserConnection{}
}

func UserToUserInformation(user model.User) UserInformation {
	return UserInformation{}
}

func UserConnectionToUser(user UserConnection) model.User {
	return model.User{}
}

func UserInformationToUser(user UserInformation) model.User {
	return model.User{}
}
