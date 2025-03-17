package shoppinglist

import "github.com/google/uuid"

type ShoppinglistRepository interface {
	Insert(shoppinglistItems []ShoppinglistItem) ([]ShoppinglistItem, error)
	FindAllForUserIds(userId uuid.UUID) ([]ShoppinglistItem, error)
	FindAllForGroupIds(groupId uuid.UUID) ([]ShoppinglistItem, error)
}
