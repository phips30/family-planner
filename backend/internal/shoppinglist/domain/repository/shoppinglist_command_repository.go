package repository

import (
	"family-planner/backend/internal/shoppinglist/domain/entity"
)

type ShoppinglistCommandRepository interface {
	Insert(shoppinglistItems []entity.ShoppinglistItem) ([]string, error)
	Delete(itemId string) error
	Update(shoppingItem entity.ShoppinglistItem) error
}
