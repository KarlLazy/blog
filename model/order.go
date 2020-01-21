package model

// Order order
type Order struct {
	OrderID      string     `bson:"orderId" json:"orderId"`
	OrderStatus  string     `bson:"orderStatus" json:"orderStatus"`
	OrderContent []*Content `bson:"orderContent" json:"orderContent"`
	Amount       float64    `bson:"amount" json:"amount"`
	ImageURL     string     `bson:"imageURL,omitempty" json:"imageURL,omitempty"`

	StoreName string   `bson:"storeName" json:"storeName"`
	StoreNum  string   `bson:"storeNum" json:"storeNum"`
	Comment   *Comment `bson:"comment,omitempty" json:"comment,omitempty"`
	Refund    *Refund  `bson:"refund,omitempty" json:"refund,omitempty"`
}

// Content content
type Content struct {
	ComName string `bson:"comName" json:"comName"`
	Number  int    `bson:"number" json:"storeName"`
}

// Comment comment
type Comment struct {
	Content string `bson:"content" json:"content"`
}

// Refund refund
type Refund struct {
	Amount int `bson:"amount" json:"amount"`
}
