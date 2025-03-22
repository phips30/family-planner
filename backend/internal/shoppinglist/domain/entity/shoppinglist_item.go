package entity

import (
	"errors"
	"family-planner/backend/internal/shoppinglist/domain"
	"time"

	"github.com/google/uuid"
)

type ShoppinglistItemCreator struct {
	Id    uuid.UUID
	Name  string
	Email string
}

type ShoppinglistItem struct {
	Id      string
	Name    string
	AddedAt time.Time
	User    ShoppinglistItemCreator
	GroupId uuid.UUID
	Bought  bool
}

func NewShoppinglistItem(name string, addedAt time.Time, user ShoppinglistItemCreator, groupId uuid.UUID, bought bool) (*ShoppinglistItem, error) {
	if addedAt == (time.Time{}) {
		addedAt = time.Now()
	}
	if name == "" {
		return nil, errors.New("no name provided")
	}
	if user == (ShoppinglistItemCreator{}) {
		return nil, errors.New("no user provided")
	}
	return &ShoppinglistItem{
		Name:    name,
		AddedAt: addedAt,
		User:    user,
		GroupId: groupId,
		Bought:  bought,
	}, nil
}

func NewShoppinglistItems(items []domain.ShoppinglistDto, userData []ShoppinglistItemCreator) *[]ShoppinglistItem {
	var shoppinglistItems []ShoppinglistItem
	for _, item := range items {
		var itemCreator ShoppinglistItemCreator
		for _, user := range userData {
			if user.Id == item.UserId {
				itemCreator = user
				break
			}
		}
		// TODO: What to do if user does not exist
		shoppinglistitem, err := NewShoppinglistItem(item.Name, item.AddedAt, itemCreator, item.GroupId, item.Bought)
		if err == nil {
			// Just omit wrong items for now
			shoppinglistItems = append(shoppinglistItems, *shoppinglistitem)
		}
	}
	return &shoppinglistItems
}

func FromExistingItems(items []domain.ShoppinglistDto, userData []ShoppinglistItemCreator) *[]ShoppinglistItem {
	var shoppinglistItems []ShoppinglistItem
	for _, item := range items {
		var itemCreator ShoppinglistItemCreator
		for _, user := range userData {
			if user.Id == item.UserId {
				itemCreator = user
				break
			}
		}
		// TODO: What to do if user does not exist
		shoppinglistitem := ShoppinglistItem{
			Id:      item.Id,
			Name:    item.Name,
			AddedAt: item.AddedAt,
			User:    itemCreator,
			GroupId: item.GroupId,
			Bought:  item.Bought,
		}
		shoppinglistItems = append(shoppinglistItems, shoppinglistitem)
	}

	return &shoppinglistItems
}
