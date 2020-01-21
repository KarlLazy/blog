package mongo

import (
	"blog/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type orderCollection struct {
	Name string
}

// OrderCol order collection
var OrderCol = orderCollection{"order"}

var ctx = context.TODO()

func (col *orderCollection) Add(order *model.Order) error {
	_, err := db.Collection(col.Name).InsertOne(context.TODO(), order)

	return err
}

func (col *orderCollection) Update(order *model.Order) error {
	cond := bson.M{
		"orderId": order.OrderID,
	}

	update := bson.M{
		"OrderStatus": order.OrderStatus,
	}

	_, err := db.Collection(col.Name).UpdateOne(context.TODO(), cond, update)

	return err
}

func (col *orderCollection) Find(orderID string) (order *model.Order, err error) {
	cond := bson.M{
		"orderId": orderID,
	}

	err = db.Collection(col.Name).FindOne(context.TODO(), cond).Decode(&order)

	return order, err
}

func (col *orderCollection) Delete(orderID string) error {
	cond := bson.M{
		"orderId": orderID,
	}

	_, err := db.Collection(col.Name).DeleteOne(context.TODO(), cond)

	return err
}

func (col *orderCollection) FindOrderByKey(index, size int, key string) (orders []model.Order, err error) {
	match := bson.M{}

	switch key {
	case model.KeyAllOrder:
	case model.KeyComment:
		match["comment"] = bson.M{
			"$exists": true,
		}
	case model.KeyRefund:
		match["refund"] = bson.M{
			"$exists": true,
		}
	}

	pipe := []bson.M{
		{"$match": match},
	}

	sort := bson.M{"$sort": bson.M{"orderId": 1}}

	skip := bson.M{"$skip": index}

	limit := bson.M{"$limit": size}

	pipe = append(pipe, sort, skip, limit)

	orders = make([]model.Order, 0)
	cur, err := db.Collection(col.Name).Aggregate(ctx, pipe)
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &orders); err != nil {
		return nil, err
	}

	// for cur.Next(ctx) {
	// 	order := model.Order{}

	// 	if err := cur.Decode(&order); err != nil {
	// 		return nil, err
	// 	}
	// 	orders = append(orders, order)
	// }

	return orders, err
}

func (col *orderCollection) Count(status string) (int64, error) {
	cond := bson.M{
		"orderStatus": status,
	}

	count, err := db.Collection(col.Name).CountDocuments(context.TODO(), cond)

	return count, err
}
