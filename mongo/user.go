package mongo

import (
	"context"
"blog/model"
)

type userCollection struct {
	Name string
}

// UserCol user collection
var UserCol = userCollection{"user"}

func (col *userCollection) AddUser(user *model.User) {
	_, _ = db.Collection(col.Name).InsertOne(context.TODO(), user)
}
