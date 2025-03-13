package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
}

func NewUser(name string, email string) (*User, error) {
	user := User{
		Id:        uuid.New(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}
	if err := user.validate(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) validate() error {
	if u.Name == "" {
		return errors.New("name cannot be empty")
	}
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}

	return nil
}
