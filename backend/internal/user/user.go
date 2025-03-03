package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	Name      string
	DeviceId  string
	CreatedAt time.Time
}

func NewUser(name string, deviceId string) (*User, error) {
	user := User{
		Id:        uuid.New(),
		Name:      name,
		DeviceId:  deviceId,
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
	if u.DeviceId == "" {
		return errors.New("deviceId cannot be empty")
	}
	if user, err := UserRepository.FindByNameAndDeviceId(nil, u.Name, u.DeviceId); err != nil && user != nil {
		return errors.New("user already exists")
	}

	return nil
}
