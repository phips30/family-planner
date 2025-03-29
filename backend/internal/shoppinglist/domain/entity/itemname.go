package entity

import "errors"

type ItemName string

func NewItemName(name string) (ItemName, error) {
	if name == "" {
		return "", errors.New("item name cannot be empty")
	}
	itemName := ItemName(name)
	return itemName, nil
}
