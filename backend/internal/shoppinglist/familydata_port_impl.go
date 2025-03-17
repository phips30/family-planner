package shoppinglist

import (
	"github.com/google/uuid"
)

type FamilyDataPortImpl struct{}

func (f *FamilyDataPortImpl) GetGroupData(userID uuid.UUID) (*GroupDataDto, error) {
	return nil, nil
}
