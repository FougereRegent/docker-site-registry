package model

import (
	"container/list"
	"fmt"
	"user-service/utils"
)

type NotFoundError struct {
	ErrorMessage string
	Table        string
	Val          string
}

func initList() *list.List {
	listSchema := list.New()
	listSchema.PushBack(&User{})
	return listSchema
}

func InitDataBase() {
	list := initList()
	db := utils.GetConnector()

	for val := list.Front(); val != nil; val = val.Next() {
		db.AutoMigrate(&val.Value)
	}
}

func (e *NotFoundError) Error() string {
	message := fmt.Sprintf("ErrorMessage : %s\nTable : %s\nValeur %s\n", e.ErrorMessage, e.Table, e.Val)
	return message
}
