package groupmember

import (
	"context"
	"fmt"

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

func (gm *GroupMemberRepositoryImpl) Save(groupMember *GroupMember) (*GroupMember, error) {
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
