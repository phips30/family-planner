package postgres

import (
	"context"
	"family-planner/backend/internal/family/domain/entity"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupMemberRepositoryImpl struct {
	ctx    context.Context
	dbpool *pgxpool.Pool
}

func NewGroupMemberRepositoryImpl(ctx context.Context, dbpool *pgxpool.Pool) *GroupMemberRepositoryImpl {
	return &GroupMemberRepositoryImpl{ctx: ctx, dbpool: dbpool}
}

func (gm *GroupMemberRepositoryImpl) Save(groupMember *entity.GroupMember) (*entity.GroupMember, error) {
	query := `INSERT INTO public.group_member (group_id, user_id, created_at) VALUES (@groupId, @userId, @createdAt)`
	args := pgx.NamedArgs{
		"groupId":   groupMember.Group.Id,
		"userId":    groupMember.User.Id,
		"createdAt": groupMember.CreatedAt,
	}
	_, err := gm.dbpool.Exec(gm.ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to insert row: %w", err)
	}

	return groupMember, nil
}

func (gm *GroupMemberRepositoryImpl) FindMembers(groupId uuid.UUID) ([]*entity.User, error) {
	query := `select u.* from public.group_member gm
				join public.user u on u.id = gm.user_id
				where gm.group_id = @groupId`
	args := pgx.NamedArgs{
		"groupId": groupId,
	}

	rows, err := gm.dbpool.Query(gm.ctx, query, args)
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
