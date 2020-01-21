package mongo

import (
	"blog/config"
	"blog/model"
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	fmt.Printf("url: %s\n", config.Config.User.Name)

	fmt.Printf("config: %s\n", config.Config.Mongo.Name)

	user := &model.User{
		UserName: "karl",
		Password: "pwd",
	}

	UserCol.AddUser(user)
}
