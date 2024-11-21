package model

type Role struct {
	RoleID          int64
	RoleName        string
	RoleDescription string
}

func NewRole(roleID int64, roleName string, roleDescription string) *Role {
	return &Role{
		RoleID:          roleID,
		RoleName:        roleName,
		RoleDescription: roleDescription,
	}
}

func (r *Role) GetID() int64 {
	return r.RoleID
}

func (r *Role) GetName() string {
	return r.RoleName
}

func (r *Role) GetDescription() string {
	return r.RoleDescription
}
