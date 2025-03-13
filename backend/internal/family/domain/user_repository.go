package domain

type UserRepository interface {
	Create(user *User) (*User, error)
	FindByEmail(email string) (*User, error)
}
