package domain

type GroupMemberService struct {
	repository GroupMemberRepository
}

func NewGroupMemberService(repository GroupMemberRepository) *GroupMemberService {
	return &GroupMemberService{
		repository: repository,
	}
}

func (g *GroupMemberService) AddMember(group Group, user User) (*GroupMember, error) {
	newGroupMember := NewGroupMember(group, user)
	return g.repository.Save(&newGroupMember)
}
