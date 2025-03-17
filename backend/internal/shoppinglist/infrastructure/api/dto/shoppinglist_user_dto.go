package dto

import (
	"family-planner/backend/internal/shoppinglist/domain/entity"
	"time"

	"github.com/google/uuid"
)

type ShoppinglistUserItemDto struct {
	Name    string    `json:"name"`
	UserId  uuid.UUID `json:"userId"`
	AddedAt time.Time `json:"addedAt"`
	Bought  bool      `json:"bought"`
}

func (itemdto *ShoppinglistUserItemDto) MapToDomainObject() *entity.ShoppinglistItem {
	return entity.NewShoppinglistItem(
		itemdto.Name,
		itemdto.AddedAt,
		itemdto.UserId,
		uuid.Nil,
		itemdto.Bought)
}

func (itemdto *ShoppinglistUserItemDto) MapFromDomainObject(shoppinglistItem entity.ShoppinglistItem) *ShoppinglistUserItemDto {
	return &ShoppinglistUserItemDto{
		Name:    shoppinglistItem.Name,
		UserId:  shoppinglistItem.UserId,
		AddedAt: shoppinglistItem.AddedAt,
		Bought:  shoppinglistItem.Bought,
	}
}
