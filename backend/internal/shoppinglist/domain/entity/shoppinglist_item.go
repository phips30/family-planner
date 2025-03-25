package entity

import (
	"errors"
	"family-planner/backend/internal/shoppinglist/domain"
	"time"
)

type ShoppinglistItemCreator struct {
	Id    UserId
	Name  string
	Email string
}

type ShoppinglistItem struct {
	Id      string
	Name    ItemName
	AddedAt time.Time
	User    ShoppinglistItemCreator
	GroupId GroupId
	Bought  bool
}

func (s *ShoppinglistItem) SetBought(bought bool) {
	s.Bought = bought
}

func (s *ShoppinglistItem) SetName(itemName ItemName) {
	s.Name = itemName
}

func NewShoppinglistItem(name ItemName, addedAt time.Time, user ShoppinglistItemCreator, groupId GroupId, bought bool) (*ShoppinglistItem, error) {
	if addedAt == (time.Time{}) {
		addedAt = time.Now()
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
			if user.Id.UUID == item.UserId {
				itemCreator = user
				break
			}
		}
		// TODO: What to do if user does not exist
		itemName, err := NewItemName(item.Name)
		if err != nil {
			break
		}

		shoppinglistitem, err := NewShoppinglistItem(*itemName, item.AddedAt, itemCreator, *NewGroupId(item.GroupId), item.Bought)
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
			if user.Id.UUID == item.UserId {
				itemCreator = user
				break
			}
		}
		shoppinglistItems = append(shoppinglistItems, *FromExistingItem(item, itemCreator))
	}

	return &shoppinglistItems
}

func FromExistingItem(item domain.ShoppinglistDto, itemCreator ShoppinglistItemCreator) *ShoppinglistItem {
	itemName, err := NewItemName(item.Name)
	if err != nil {
		return nil
	}

	// TODO: What to do if user does not exist
	return &ShoppinglistItem{
		Id:      item.Id,
		Name:    *itemName,
		AddedAt: item.AddedAt,
		User:    itemCreator,
		GroupId: *NewGroupId(item.GroupId),
		Bought:  item.Bought,
	}
}
