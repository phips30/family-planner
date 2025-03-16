package shoppinglist

import "github.com/google/uuid"

type ShoppinglistRepository interface {
	Insert(shoppinglistItems []ShoppinglistItem) ([]ShoppinglistItem, error)
	FindAllForUserIds(userIds []uuid.UUID) ([]ShoppinglistItem, error)
}
