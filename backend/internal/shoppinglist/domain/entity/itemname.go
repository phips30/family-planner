package entity

import "errors"

type ItemName string

func (i ItemName) ToString() string {
	return string(i)
}

func NewItemName(name string) (*ItemName, error) {
	if name == "" {
		return nil, errors.New("no name provided")
	}
	itemName := ItemName(name)
	return &itemName, nil
}
