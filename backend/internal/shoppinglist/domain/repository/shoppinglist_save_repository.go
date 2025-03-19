package repository

import (
	"family-planner/backend/internal/shoppinglist/common/dto"
)

type ShoppinglistSaveRepository interface {
	Insert(shoppinglistItems []dto.ShoppinglistItemRequestDto) ([]dto.ShoppinglistItemRequestDto, error)
}
