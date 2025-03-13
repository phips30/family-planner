package repository

import "family-planner/backend/internal/family/domain/entity"

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}
