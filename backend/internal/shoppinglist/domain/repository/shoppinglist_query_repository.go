package repository

import (
	"family-planner/backend/internal/shoppinglist/common/dto"

	"github.com/google/uuid"
)

type ShoppinglistQueryRepository interface {
	FindAll(id uuid.UUID) ([]dto.ShoppinglistItemRequestDto, error)
}
