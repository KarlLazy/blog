package mongo

import (
	"blog/model"
	"testing"
)

func TestAddUser(t *testing.T) {
	user := &model.User{
		UserName: "karl",
		Password: "pwd",
	}

	UserCol.AddUser(user)
}
