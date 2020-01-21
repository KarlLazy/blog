package controller

import (
	"blog/model"
	"blog/mongo"
	"log"
)

type order struct{}

// Order order
var Order order

func (o *order) FindOrder(index, limit int, key string) []model.Order {
	orders, err := mongo.OrderCol.FindOrderByKey(index, limit, key)

	if err != nil {
		log.Printf("error: %v", err)
		return nil
	}

	return orders
}
