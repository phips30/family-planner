package repository

import (
	"family-planner/backend/internal/shoppinglist/domain/entity"
)

type ShoppinglistSaveRepository interface {
	Insert(shoppinglistItems []entity.ShoppinglistItem) ([]entity.ShoppinglistItem, error)
}
