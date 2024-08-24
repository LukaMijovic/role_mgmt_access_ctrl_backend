package model

import "time"

type TempAccessRight struct {
	accessID        int64
	startAccessDate time.Time
	endAccessDate   time.Time
	userID          int64
}

func (t *TempAccessRight) GetID() int64 {
	return t.accessID
}

func (t *TempAccessRight) GetStartAccessDate() time.Time {
	return t.startAccessDate
}

func (t *TempAccessRight) GetEndAccessDate() time.Time {
	return t.endAccessDate
}

func (t *TempAccessRight) GetUserID() int64 {
	return t.userID
}
