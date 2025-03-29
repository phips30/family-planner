package postgres

import (
	"context"
	"family-planner/backend/internal/common/models"
	"log"

	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserDataAdapter struct {
	ctx    context.Context
	dbpool *pgxpool.Pool
}

func NewUserDataAdapter(ctx context.Context, dbpool *pgxpool.Pool) *UserDataAdapter {
	return &UserDataAdapter{ctx: ctx, dbpool: dbpool}
}

// TODO: Change return type
func (u *UserDataAdapter) GetUserData(userIDs []uuid.UUID) ([]models.UserDto, error) {
	query := `SELECT u.id, u.name, u.email FROM public.user u 
	where u.id = ANY($1)`

	var users []models.UserDto
	rows, err := u.dbpool.Query(u.ctx, query, userIDs)

	for rows.Next() {
		var user models.UserDto
		err := rows.Scan(&user.Id, &user.Name, &user.Email)

		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}

	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("unable to query user table: %w", err)
	}
	return users, nil
}
