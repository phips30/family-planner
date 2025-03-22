package dto

import (
	"time"

	"github.com/google/uuid"
)

type ShoppingtItemCreatorResponseDto struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type ShoppinglistItemResponseDto struct {
	Id      string                          `json:"id"`
	Name    string                          `json:"name"`
	User    ShoppingtItemCreatorResponseDto `json:"user"`
	AddedAt time.Time                       `json:"addedAt"`
	Bought  bool                            `json:"bought"`
}
