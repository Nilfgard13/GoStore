package app

import "github.com/Nilfgard13/GOSTORE/app/model"

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: model.User{}},
		{Model: model.Address{}},
		{Model: model.Product{}},
		{Model: model.ProductImage{}},
		{Model: model.Section{}},
		{Model: model.Category{}},
		{Model: model.Order{}},
		{Model: model.OrderItem{}},
		{Model: model.OrderCustomer{}},
		{Model: model.Payment{}},
		{Model: model.Shipment{}},
	}
}
