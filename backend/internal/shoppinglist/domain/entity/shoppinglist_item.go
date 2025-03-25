package entity

import (
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
	UserId  UserId
	GroupId GroupId
	Bought  bool
}

func (s *ShoppinglistItem) SetBought(bought bool) {
	s.Bought = bought
}

func (s *ShoppinglistItem) SetName(itemName ItemName) {
	s.Name = itemName
}

func NewShoppinglistItem(name ItemName, addedAt time.Time, user UserId, groupId GroupId, bought bool) (*ShoppinglistItem, error) {
	if addedAt == (time.Time{}) {
		addedAt = time.Now()
	}
	return &ShoppinglistItem{
		Name:    name,
		AddedAt: addedAt,
		UserId:  user,
		GroupId: groupId,
		Bought:  bought,
	}, nil
}

func NewShoppinglistItems(items []domain.ShoppinglistDto) *[]ShoppinglistItem {
	var shoppinglistItems []ShoppinglistItem
	for _, item := range items {
		itemName, err := NewItemName(item.Name)
		if err != nil {
			break
		}

		userId, err := NewUserId(item.UserId)
		if err != nil {
			break
		}

		shoppinglistitem, err := NewShoppinglistItem(*itemName, item.AddedAt, *userId, *NewGroupId(item.GroupId), item.Bought)
		if err == nil {
			// Just omit wrong items for now
			shoppinglistItems = append(shoppinglistItems, *shoppinglistitem)
		}
	}
	return &shoppinglistItems
}

func FromExistingItems(items []domain.ShoppinglistDto) *[]ShoppinglistItem {
	var shoppinglistItems []ShoppinglistItem
	for _, item := range items {
		shoppinglistItems = append(shoppinglistItems, *FromExistingItem(item))
	}

	return &shoppinglistItems
}

func FromExistingItem(item domain.ShoppinglistDto) *ShoppinglistItem {
	itemName, err := NewItemName(item.Name)
	if err != nil {
		return nil
	}

	userId, err := NewUserId(item.UserId)
	if err != nil {
		return nil
	}

	return &ShoppinglistItem{
		Id:      item.Id,
		Name:    *itemName,
		AddedAt: item.AddedAt,
		UserId:  *userId,
		GroupId: *NewGroupId(item.GroupId),
		Bought:  item.Bought,
	}
}
