package shoppinglist

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ShoppinglistService struct {
	repository     ShoppinglistRepository
	familyDataPort FamilyDataPort
}

var (
	errRetrievingGroupData = errors.New("error getting group information for user")
)

func NewShoppinglistService(repository ShoppinglistRepository, familyDataPort FamilyDataPort) *ShoppinglistService {
	return &ShoppinglistService{repository: repository, familyDataPort: familyDataPort}
}

func (s *ShoppinglistService) LoadShoppingListForUserid(userid uuid.UUID) ([]ShoppinglistItem, error) {
	groupData, err := s.familyDataPort.GetGroupData(userid)
	if err != nil {
		return nil, errRetrievingGroupData
	}
	if groupData == nil {
		return s.repository.FindAllForUserIds([]uuid.UUID{userid})
	} else {
		var userIds []uuid.UUID
		for _, member := range groupData.Members {
			userIds = append(userIds, member.UserId)
		}
		return s.repository.FindAllForUserIds(userIds)
	}
}

func (s *ShoppinglistService) SaveShoppingList(shoppinglist []ShoppinglistItem) ([]ShoppinglistItem, error) {
	validation_error := s.validate(shoppinglist)
	// TODO: Just aggregate the errors and log them instead of not continuing
	if validation_error != nil {
		return nil, fmt.Errorf("%s", validation_error.Error())
	}
	return s.repository.Insert(shoppinglist)
}

func (s *ShoppinglistService) validate(shoppinglistItems []ShoppinglistItem) error {
	// TODO: Check if item does not already exist for the group for a shopping trip
	for _, shoppinglistItem := range shoppinglistItems {
		if shoppinglistItem.Name == "" {
			return errors.New("no name provided")
		}
		if shoppinglistItem.AddedBy == "" {
			return errors.New("no user provided")
		}
		if shoppinglistItem.AddedAt == (time.Time{}) {
			shoppinglistItem.AddedAt = time.Now()
		}
	}
	return nil
}
