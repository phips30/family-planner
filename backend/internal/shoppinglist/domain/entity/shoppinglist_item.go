package entity

import (
	"family-planner/backend/internal/shoppinglist/common/dto"
	"time"

	"github.com/google/uuid"
)

type ShoppinglistItemCreator struct {
	Id    uuid.UUID
	Name  string
	Email string
}

type ShoppinglistItem struct {
	Name    string
	AddedAt time.Time
	User    ShoppinglistItemCreator
	GroupId uuid.UUID
	Bought  bool
}

func NewShoppinglistItem(name string, addedAt time.Time, user ShoppinglistItemCreator, groupId uuid.UUID, bought bool) *ShoppinglistItem {
	if addedAt == (time.Time{}) {
		addedAt = time.Now()
	}
	return &ShoppinglistItem{
		Name:    name,
		AddedAt: addedAt,
		User:    user,
		GroupId: groupId,
		Bought:  bought,
	}
}

func FromExistingItems(items []dto.ShoppinglistItemRequestDto, userData []ShoppinglistItemCreator) *[]ShoppinglistItem {
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
