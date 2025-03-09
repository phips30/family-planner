package shoppinglist

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type ShoppinglistItem struct {
	Name    string
	AddedAt time.Time
	AddedBy uuid.UUID
	Bought  bool
	Sorter  int
}

func NewShoppinglistItem(name string, addedAt time.Time, addedBy uuid.UUID, bought bool, sorter int) (*ShoppinglistItem, error) {
	shoppinglistItem := ShoppinglistItem{
		Name:    name,
		AddedAt: addedAt,
		AddedBy: addedBy,
		Bought:  bought,
		Sorter:  sorter,
	}
	if err := shoppinglistItem.validate(); err != nil {
		return nil, err
	}

	return &shoppinglistItem, nil
}

func (s *ShoppinglistItem) validate() error {
	if s.Name == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}
