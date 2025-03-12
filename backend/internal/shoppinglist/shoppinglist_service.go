package shoppinglist

import (
	"errors"
	"fmt"
	"time"
)

type ShoppinglistService struct {
	repository ShoppinglistRepositoryImpl
}

func NewShoppinglistService(repository ShoppinglistRepositoryImpl) *ShoppinglistService {
	return &ShoppinglistService{repository: repository}
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
