package dto

import (
	"family-planner/backend/internal/shoppinglist/domain/entity"
	"time"

	"github.com/google/uuid"
)

type ShoppinglistGroupItemDto struct {
	Name    string    `json:"name"`
	UserId  uuid.UUID `json:"userId"`
	GroupId uuid.UUID `json:"groupId"`
	AddedAt time.Time `json:"addedAt"`
	Bought  bool      `json:"bought"`
}

func (itemdto *ShoppinglistGroupItemDto) MapToDomainObject() *entity.ShoppinglistItem {
	return entity.NewShoppinglistItem(
		itemdto.Name,
		itemdto.AddedAt,
		itemdto.UserId,
		itemdto.GroupId,
		itemdto.Bought)
}

func (itemdto *ShoppinglistGroupItemDto) MapFromDomainObject(shoppinglistItem entity.ShoppinglistItem) *ShoppinglistGroupItemDto {
	return &ShoppinglistGroupItemDto{
		Name:    shoppinglistItem.Name,
		UserId:  shoppinglistItem.UserId,
		GroupId: shoppinglistItem.GroupId,
		AddedAt: shoppinglistItem.AddedAt,
		Bought:  shoppinglistItem.Bought,
	}
}
