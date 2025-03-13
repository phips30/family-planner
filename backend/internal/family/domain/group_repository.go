package domain

type GroupRepository interface {
	Save(group *Group) (*Group, error)
	FindGroupForUser(email string) *Group
}
