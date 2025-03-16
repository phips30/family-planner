package shoppinglist

import (
	"github.com/google/uuid"
)

type GroupDataDto struct {
	Name    string
	Members []GroupMemerDataDto
}

type GroupMemerDataDto struct {
	UserId uuid.UUID
	Name   string
	Email  string
}

type FamilyDataPort interface {
	GetGroupData(userID uuid.UUID) (*GroupDataDto, error)
}
