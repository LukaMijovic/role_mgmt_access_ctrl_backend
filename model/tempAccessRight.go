package model

import "time"

type TempAccessRight struct {
	accessID        int64
	accessName      string
	startAccessDate time.Time
	EndAccessDate   time.Time `binding:required`
	UserID          int64     `binding:required`
}

func (t *TempAccessRight) GetID() int64 {
	return t.accessID
}

func (t *TempAccessRight) SetAccessId(accessId int64) {
	t.accessID = accessId
}

func (t *TempAccessRight) GetStartAccessDate() time.Time {
	return t.startAccessDate
}

func (t *TempAccessRight) GetEndAccessDate() time.Time {
	return t.EndAccessDate
}

func (t *TempAccessRight) GetUserID() int64 {
	return t.UserID
}

func (t *TempAccessRight) GetAccessName() string {
	return t.accessName
}

func (t *TempAccessRight) SetAccessName(accessName string) {
	t.accessName = accessName
}
