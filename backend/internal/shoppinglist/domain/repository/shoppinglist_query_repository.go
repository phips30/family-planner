package repository

import (
	"family-planner/backend/internal/shoppinglist/domain/entity"

	"github.com/google/uuid"
)

type ShoppinglistQueryRepository interface {
	FindAll(id uuid.UUID) ([]entity.ShoppinglistItem, error)
}
