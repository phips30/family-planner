package dto

import (
	"time"

	"github.com/google/uuid"
)

type ShoppingtItemCreator struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type ShoppinglistItemResponse struct {
	Name    string               `json:"name"`
	User    ShoppingtItemCreator `json:"user"`
	AddedAt time.Time            `json:"addedAt"`
	Bought  bool                 `json:"bought"`
}
