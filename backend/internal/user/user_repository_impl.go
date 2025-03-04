package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryImpl struct {
	dbpool *pgxpool.Pool
}

func NewUserRepositoryImpl(dbpool *pgxpool.Pool) *UserRepositoryImpl {
	return &UserRepositoryImpl{dbpool: dbpool}
}

func (u *UserRepositoryImpl) Create(user *User) (*User, error) {
	return nil, nil
}

func (u *UserRepositoryImpl) FindByNameAndDeviceId(name string, deviceId string) (*User, error) {
	fmt.Printf("Getting user data")

	var userInDb User
	query := `SELECT * FROM user LIMIT 10`

	fmt.Print("Lets see if an error occurs")
	rows, err := u.dbpool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}
	defer rows.Close()

	fmt.Print("There was no error unitl here")
	users := []User{}
	for rows.Next() {
		fmt.Print("Lets see if we can append")
		user := User{}
		err := rows.Scan(&user.Id, &user.Name, &user.DeviceId)
		fmt.Print("Lets see if we can append here")

		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		fmt.Print("Lets see if we can append here 2")
		users = append(users, user)
	}
	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println("No rows found")
		} else {
			fmt.Println("An error occured")
			fmt.Println(err)
		}
	}

	fmt.Printf("%s", userInDb.Name)

	return &userInDb, nil
}
