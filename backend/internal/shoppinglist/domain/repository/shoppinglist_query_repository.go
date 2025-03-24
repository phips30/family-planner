package repository

import (
	"family-planner/backend/internal/shoppinglist/domain"

	"github.com/google/uuid"
)

type ShoppinglistQueryRepository interface {
	FindAllByGroupId(groupId uuid.UUID) ([]domain.ShoppinglistDto, error)
	FindAllByUserId(userId uuid.UUID) ([]domain.ShoppinglistDto, error)
	FindById(id string) (*domain.ShoppinglistDto, error)
}
