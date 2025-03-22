package repository

import (
	"family-planner/backend/internal/shoppinglist/domain"

	"github.com/google/uuid"
)

type ShoppinglistQueryRepository interface {
	FindAll(id uuid.UUID) ([]domain.ShoppinglistDto, error)
}
