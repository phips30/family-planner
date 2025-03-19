package postgres

import (
	"context"
	shoppinglist "family-planner/backend/internal/shoppinglist/domain/entity"
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
func (u *UserDataAdapter) GetUserData(userIDs []uuid.UUID) ([]shoppinglist.ShoppinglistItemCreator, error) {
	query := `SELECT u.id, u.name, u.email FROM public.user u 
	where u.id = ANY($1)`

	var users []shoppinglist.ShoppinglistItemCreator
	rows, err := u.dbpool.Query(u.ctx, query, userIDs)

	for rows.Next() {
		var user shoppinglist.ShoppinglistItemCreator
		err := rows.Scan(&user.Id, &user.Name, &user.Email)

		if err != nil {
			log.Println(err)
			return nil, err
		}
		fmt.Printf("ID: %d, Name: %s\n", user.Id, user.Name)
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
