package shoppinglist

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type ShoppinglistService struct {
	repository ShoppinglistRepository
}

var (
	errRetrievingUserData  = errors.New("error getting shopping list for user")
	errRetrievingGroupData = errors.New("error getting shopping list for group")
)

func NewShoppinglistService(repository ShoppinglistRepository) *ShoppinglistService {
	return &ShoppinglistService{repository: repository}
}

func (s *ShoppinglistService) LoadShoppingListForUserId(userid uuid.UUID) ([]ShoppinglistItem, error) {
	items, err := s.repository.FindAllForUserIds(userid)
	if err != nil {
		log.Println(err.Error())
		return nil, errRetrievingUserData
	}
	return items, nil
}

func (s *ShoppinglistService) LoadShoppingListForGroupId(groupId uuid.UUID) ([]ShoppinglistItem, error) {
	items, err := s.repository.FindAllForGroupIds(groupId)
	if err != nil {
		log.Println(err.Error())
		return nil, errRetrievingGroupData
	}
	return items, nil
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
		if shoppinglistItem.UserId == (uuid.UUID{}) {
			return errors.New("no user provided")
		}
		if shoppinglistItem.AddedAt == (time.Time{}) {
			shoppinglistItem.AddedAt = time.Now()
		}
	}
	return nil
}
