package user

type UserRepository interface {
	Create(user *User) (*User, error)
	FindByNameAndDeviceId(name string, deviceId string) (*User, error)
}
