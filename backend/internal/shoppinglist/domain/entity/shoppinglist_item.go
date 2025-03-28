package entity

import (
	"time"

	"github.com/google/uuid"
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

func NewShoppinglistItem(name string, addedAt time.Time, userId uuid.UUID, groupId uuid.UUID, bought bool) (*ShoppinglistItem, error) {
	itemName, err := NewItemName(name)
	if err != nil {
		return nil, err
	}

	user, err := NewUserId(userId)
	if err != nil {
		return nil, err
	}

	group := NewGroupId(groupId)

	if addedAt == (time.Time{}) {
		addedAt = time.Now()
	}
	return &ShoppinglistItem{
		Name:    *itemName,
		AddedAt: addedAt,
		UserId:  *user,
		GroupId: *group,
		Bought:  bought,
	}, nil
}

func FromExistingItem(id string, name string, addedAt time.Time, userId uuid.UUID, groupId uuid.UUID, bought bool) (*ShoppinglistItem, error) {
	item, err := NewShoppinglistItem(name, addedAt, userId, groupId, bought)
	if err != nil {
		return nil, err
	}
	item.Id = id
	return item, nil
}
