package mongo

import (
	"blog/model"
	"context"
)

type userCollection struct {
	Name string
}

// UserCol user collection
var UserCol = userCollection{"user"}

// AddUser add user
func (col *userCollection) AddUser(user *model.User) {
	_, _ = db.Collection(col.Name).InsertOne(context.TODO(), user)
}
