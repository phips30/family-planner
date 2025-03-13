package repository

import (
	"context"
	"family-planner/backend/internal/family/domain"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupRepositoryImpl struct {
	ctx    context.Context
	dbpool *pgxpool.Pool
}

func NewGroupRepositoryImpl(ctx context.Context, dbpool *pgxpool.Pool) *GroupRepositoryImpl {
	return &GroupRepositoryImpl{ctx: ctx, dbpool: dbpool}
}

func (g *GroupRepositoryImpl) Save(group *domain.Group) (*domain.Group, error) {
	query := `INSERT INTO public.group (id, name, created_by, created_at) VALUES (@id, @name, @createdBy, @createdAt)`
	args := pgx.NamedArgs{
		"id":        group.Id,
		"name":      group.Name,
		"createdBy": group.CreatedBy.Id,
		"createdAt": group.CreatedAt,
	}
	_, err := g.dbpool.Exec(g.ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to insert row: %w", err)
	}

	return group, nil
}

func (g *GroupRepositoryImpl) FindGroupForUser(requestedByEmail string) *domain.Group {
	panic("unimplemented")
}
