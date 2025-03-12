package shoppinglist

import "github.com/google/uuid"

type ShoppinglistRepository interface {
	Insert(shoppinglistItems []ShoppinglistItem) ([]ShoppinglistItem, error)
	FindAllInGroup(group uuid.UUID) ([]ShoppinglistItem, error)
}
