package model

type RoleAccessRight struct {
	accessID int64
	roleID   int64
}

func (r *RoleAccessRight) GetAccessID() int64 {
	return r.accessID
}

func (r *RoleAccessRight) GetRoleID() int64 {
	return r.roleID
}
