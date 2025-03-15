package postgres

import (
	"context"
	"family-planner/backend/internal/family/domain/aggregate"
	"family-planner/backend/internal/family/domain/entity"
	"family-planner/backend/internal/family/domain/repository"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupRepositoryImpl struct {
	ctx            context.Context
	dbpool         *pgxpool.Pool
	userRepository repository.UserRepository
}

func NewGroupRepositoryImpl(ctx context.Context, dbpool *pgxpool.Pool, userRepository repository.UserRepository) *GroupRepositoryImpl {
	return &GroupRepositoryImpl{
		ctx:            ctx,
		dbpool:         dbpool,
		userRepository: userRepository,
	}
}

func (g *GroupRepositoryImpl) Save(group *entity.Group) (*entity.Group, error) {
	query := `INSERT INTO public.group (id, name, created_by, created_at) VALUES (@id, @name, @createdBy, @createdAt)`
	fmt.Println(group.Id)
	fmt.Println(group.Name)
	fmt.Println(group.CreatedBy.Id)
	fmt.Println(group.CreatedAt)
	args := pgx.NamedArgs{
		"id":        &group.Id,
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

func (g *GroupRepositoryImpl) SaveGroupMember(group *entity.Group, groupMembers []*entity.GroupMember) error {
	query := `INSERT INTO public.group_member (group_id, user_id, created_at) VALUES (@groupId, @userId, @createdAt)`

	batch := &pgx.Batch{}
	for _, groupMember := range groupMembers {
		args := pgx.NamedArgs{
			"groupId":   group.Id,
			"userId":    groupMember.User.Id,
			"createdAt": groupMember.CreatedAt,
		}
		batch.Queue(query, args)
	}

	results := g.dbpool.SendBatch(g.ctx, batch)

	for _ = range groupMembers {
		_, err := results.Exec()
		if err != nil {
			return fmt.Errorf("unable to insert row: %w", err)
		}
	}
	return nil
}

func (g *GroupRepositoryImpl) Find(groupId uuid.UUID) (*aggregate.GroupAgg, error) {
	query := `select g.id, g.name, g.created_by, g.created_at from public.group g
				where g.id = @id`
	args := pgx.NamedArgs{
		"id": groupId,
	}

	var group entity.Group
	var createdById uuid.UUID
	err := g.dbpool.QueryRow(g.ctx, query, args).Scan(
		&group.Id,
		&group.Name,
		&createdById,
		&group.CreatedAt)
	if err == pgx.ErrNoRows {
		return nil, repository.ErrGroupNotFound
	}

	groupCreator, err := g.userRepository.FindById(createdById)
	if err != nil {
		log.Print(err.Error())
	}

	groupMembers, err := g.findMembers(group.Id)
	if err != nil {
		log.Print(err.Error())
	}

	return aggregate.FromExistingGroupAgg(&group, groupCreator, groupMembers), nil
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

func (gm *GroupRepositoryImpl) findMembers(groupId uuid.UUID) ([]*entity.GroupMember, error) {
	query := `select gm.created_at, u.* from public.group_member gm
				join public.user u on u.id = gm.user_id
				where gm.group_id = @groupId`
	args := pgx.NamedArgs{
		"groupId": groupId,
	}

	rows, err := gm.dbpool.Query(gm.ctx, query, args)
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}
	var groupMembers []*entity.GroupMember
	for rows.Next() {
		var groupMember entity.GroupMember
		var user entity.User

		err := rows.Scan(
			&groupMember.CreatedAt,
			&user.Id,
			&user.Name,
			&user.Email,
			&user.CreatedAt)
		if err != nil {
			return nil, err
		}
		groupMember.User = &user
		groupMembers = append(groupMembers, &groupMember)
	}

	return groupMembers, nil
}
