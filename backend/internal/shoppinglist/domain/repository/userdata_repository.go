package repository

import (
	"family-planner/backend/internal/common/models"

	"github.com/google/uuid"
)

type UserDataRepository interface {
	GetUserData(userIDs []uuid.UUID) ([]models.UserDto, error)
}
