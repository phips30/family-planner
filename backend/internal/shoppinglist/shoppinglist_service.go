package shoppinglist

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type ShoppinglistService struct {
	repository *ShoppinglistRepositoryImpl
}

func NewShoppinglistService(repository *ShoppinglistRepositoryImpl) *ShoppinglistService {
	return &ShoppinglistService{repository: repository}
}

func (s *ShoppinglistService) SaveShoppingList(shoppinglistItemDtos []ShoppinglistItemDto) ([]ShoppinglistItem, error) {
	var shoppinglist []ShoppinglistItem

	for _, shoppinglistItemDto := range shoppinglistItemDtos {
		shoppinglistItem := NewShoppinglistItem(
			shoppinglistItemDto.Name,
			shoppinglistItemDto.AddedAt,
			shoppinglistItemDto.AddedBy.Email,
			shoppinglistItemDto.Bought)
		validation_error := s.validate(shoppinglistItem)
		// TODO: Just aggregate the errors and log them instead of not continuing
		if validation_error != nil {
			return nil, fmt.Errorf("%s", validation_error.Error())
		}

		shoppinglist = append(shoppinglist, *shoppinglistItem)
	}
	log.Print(shoppinglist)
	return s.repository.Insert(shoppinglist)
}

func (s *ShoppinglistService) validate(shoppinglistItem *ShoppinglistItem) error {
	// TODO: Check if item does not already exist for the group
	if shoppinglistItem.Name == "" {
		return errors.New("no name provided")
	}
	if shoppinglistItem.AddedBy == "" {
		return errors.New("no user provided")
	}
	if shoppinglistItem.AddedAt == (time.Time{}) {
		shoppinglistItem.AddedAt = time.Now()
	}
	return nil
}
