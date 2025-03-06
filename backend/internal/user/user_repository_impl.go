package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryImpl struct {
	ctx    context.Context
	dbpool *pgxpool.Pool
}

func NewUserRepositoryImpl(ctx context.Context, dbpool *pgxpool.Pool) *UserRepositoryImpl {
	return &UserRepositoryImpl{ctx: ctx, dbpool: dbpool}
}

func (u *UserRepositoryImpl) Create(user *User) (*User, error) {
	query := `INSERT INTO public.user (id, name, email, created_at) VALUES (@id, @name, @email, @createdAt)`
	args := pgx.NamedArgs{
		"id":        user.Id,
		"name":      user.Name,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
	}
	_, err := u.dbpool.Exec(u.ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to insert row: %w", err)
	}

	return user, nil
}

func (u *UserRepositoryImpl) FindByEmail(email string) (*User, error) {
	var userInDb User
	query := `SELECT * FROM public.user where email = @email`
	args := pgx.NamedArgs{
		"email": email,
	}

	err := u.dbpool.QueryRow(u.ctx, query, args).Scan(&userInDb.Id, &userInDb.Name, &userInDb.Email, &userInDb.CreatedAt)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("unable to query user table: %w", err)
	}

	return &userInDb, nil
}
