package repository

import (
	"family-planner/backend/internal/shoppinglist/domain/entity"

	"github.com/google/uuid"
)

type ShoppinglistQueryRepository interface {
	FindAllByGroupId(groupId uuid.UUID) ([]entity.ShoppinglistItem, error)
	FindAllByUserId(userId uuid.UUID) ([]entity.ShoppinglistItem, error)
	FindById(id string) (*entity.ShoppinglistItem, error)
}
