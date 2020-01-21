package controller

import (
	"blog/model"
	"blog/mongo"
	"fmt"
)

type user struct{}

// User user
var User user

func (u *user) LoginGet(*model.User) {
	fmt.Println("I can do it")
	user := &model.User{
		UserName: "karl",
		Password: "pwd",
	}
	mongo.UserCol.AddUser(user)
}

func (u *user) LoginPost(*model.User) {
	fmt.Println("I can do it again")
	user := &model.User{
		UserName: "karl",
		Password: "pwd",
	}

	mongo.UserCol.AddUser(user)
}
