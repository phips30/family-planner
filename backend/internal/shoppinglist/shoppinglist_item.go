package shoppinglist

import (
	"time"
)

type ShoppinglistItem struct {
	Name    string
	AddedAt time.Time
	AddedBy string
	Bought  bool
}

func NewShoppinglistItem(name string, addedAt time.Time, addedBy string, bought bool) *ShoppinglistItem {
	return &ShoppinglistItem{
		Name:    name,
		AddedAt: addedAt,
		AddedBy: addedBy,
		Bought:  bought,
	}
}
