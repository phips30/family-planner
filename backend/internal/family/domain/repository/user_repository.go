package repository

import (
	"errors"
	"family-planner/backend/internal/family/domain/entity"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	FindById(userId uuid.UUID) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}
