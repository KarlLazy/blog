package mongo

import (
	"blog/model"
	"testing"
)

func TestAddOrder(t *testing.T) {
	content1 := &model.Content{
		ComName: "百事可乐",
		Number:  1,
	}
	content2 := &model.Content{
		ComName: "陕西腊汁肉夹馍",
		Number:  1,
	}
	content3 := &model.Content{
		ComName: "花少爷豪华擀面皮",
		Number:  1,
	}

	comment := &model.Comment{
		Content: "Are You OK?",
	}

	// refund := &model.Refund{
	// 	Amount: 25,
	// }

	content := make([]*model.Content, 0)
	content = append(content, content1, content2, content3)
	order := &model.Order{
		OrderID:      "1000001",
		OrderStatus:  model.OrderRefund,
		OrderContent: content,
		Amount:       25.22,
		ImageURL:     "",
		StoreName:    "花少爷凉皮",
		StoreNum:     "S000001",
		Comment:      comment,
	}

	err := OrderCol.Add(order)

	t.Logf("error: %v", err)
}

func TestFindOrderByKey(t *testing.T) {
	index := 0
	size := 4
	key := model.KeyAllOrder
	orders, _ := OrderCol.FindOrderByKey(index, size, key)
	t.Logf("orders: %+v", orders)
}

func TestOrderCount(t *testing.T) {
	status := model.OrderSuccess

	count, err := OrderCol.Count(status)

	if err != nil {
		t.Errorf("error: %v", err)
	} else {
		t.Logf("count: %d", count)
	}
}
