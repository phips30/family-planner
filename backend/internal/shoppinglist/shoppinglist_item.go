package shoppinglist

import (
	"time"

	"github.com/google/uuid"
)

type ShoppinglistItem struct {
	Name    string
	AddedAt time.Time
	UserId  uuid.UUID
	Bought  bool
}

func NewShoppinglistItem(name string, addedAt time.Time, userId uuid.UUID, bought bool) *ShoppinglistItem {
	if addedAt == (time.Time{}) {
		addedAt = time.Now()
	}
	return &ShoppinglistItem{
		Name:    name,
		AddedAt: addedAt,
		UserId:  userId,
		Bought:  bought,
	}
}
