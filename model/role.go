package model

type Role struct {
	roleID          int64
	roleName        string
	roleDescription string
}

func (r *Role) GetID() int64 {
	return r.roleID
}

func (r *Role) GetName() string {
	return r.roleName
}

func (r *Role) GetDescription() string {
	return r.roleDescription
}
