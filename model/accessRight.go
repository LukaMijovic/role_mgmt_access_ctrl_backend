package model

type AccessRight struct {
	accessID   int64
	accessName string
}

func (a *AccessRight) GetID() int64 {
	return a.accessID
}

func (a *AccessRight) GetName() string {
	return a.accessName
}
