package postgres

import (
	"context"
	"family-planner/backend/internal/family/domain/entity"
	"fmt"

	"github.com/google/uuid"
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

func (u *UserRepositoryImpl) Create(user *entity.User) (*entity.User, error) {
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

func (u *UserRepositoryImpl) FindById(userId uuid.UUID) (*entity.User, error) {
	query := `SELECT u.* FROM public.user u 
				where u.id = @id`
	args := pgx.NamedArgs{
		"id": userId,
	}
	return u.findAndMapToUser(query, args)
}

func (u *UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	query := `SELECT u.* FROM public.user u 
				where u.email = @email`
	args := pgx.NamedArgs{
		"email": email,
	}
	return u.findAndMapToUser(query, args)
}

func (u *UserRepositoryImpl) findAndMapToUser(query string, args pgx.NamedArgs) (*entity.User, error) {
	var user entity.User
	err := u.dbpool.QueryRow(u.ctx, query, args).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)

	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("unable to query user table: %w", err)
	}
	return &user, nil
}
