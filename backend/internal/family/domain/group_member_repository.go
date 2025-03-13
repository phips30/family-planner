package domain

type GroupMemberRepository interface {
	Save(groupMember *GroupMember) (*GroupMember, error)
}
