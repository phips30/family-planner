package port

import (
	"family-planner/backend/internal/shoppinglist/domain/entity"

	"github.com/google/uuid"
)

type UserDataPort interface {
	GetUserData(userIDs []uuid.UUID) ([]entity.ShoppinglistItemCreator, error)
}
