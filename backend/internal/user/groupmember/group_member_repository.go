package groupmember

type GroupMemberRepository interface {
	Save(groupMember *GroupMember) (*GroupMember, error)
}
