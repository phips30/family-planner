package postgres

import (
	"context"
	"family-planner/backend/internal/family/domain/aggregate"
	"family-planner/backend/internal/family/domain/entity"
	"family-planner/backend/internal/family/domain/repository"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupRepositoryImpl struct {
	ctx                   context.Context
	dbpool                *pgxpool.Pool
	groupMemberRepository GroupMemberRepositoryImpl
}

func NewGroupRepositoryImpl(ctx context.Context, dbpool *pgxpool.Pool) *GroupRepositoryImpl {
	return &GroupRepositoryImpl{
		ctx:                   ctx,
		dbpool:                dbpool,
		groupMemberRepository: *NewGroupMemberRepositoryImpl(ctx, dbpool)}
}

func (g *GroupRepositoryImpl) Save(group *entity.Group) (*entity.Group, error) {
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

func (g *GroupRepositoryImpl) Find(groupId uuid.UUID) (*aggregate.GroupAgg, error) {
	var group entity.Group
	query := `select g.* from public.group g
				where g.id = @id`
	args := pgx.NamedArgs{
		"id": groupId,
	}

	err := g.dbpool.QueryRow(g.ctx, query, args).Scan(&group.Id, &group.Name, &group.CreatedBy, &group.CreatedAt)
	if err == pgx.ErrNoRows {
		return nil, repository.ErrGroupNotFound
	}

	fmt.Println(group.Id)
	groupMembers, err := g.groupMemberRepository.FindMembers(group.Id)
	if err != nil {
		return nil, err
	}

	groupAgg := aggregate.GroupAgg{
		Group:        &group,
		GroupMembers: groupMembers,
	}

	return &groupAgg, nil
}

func (g *GroupRepositoryImpl) FindGroupForUser(userEmail string) *entity.Group {
	var group entity.Group
	query := `select g.* from public.group_member gm
				join public.group g on g.id = gm.group_id
				join public.user u on u.id = gm.user_id
				where u.email = @email`
	args := pgx.NamedArgs{
		"email": userEmail,
	}

	err := g.dbpool.QueryRow(g.ctx, query, args).Scan(&group.Id, &group.Name, &group.CreatedBy, &group.CreatedAt)
	if err == pgx.ErrNoRows {
		return nil
	}

	return &group
}
