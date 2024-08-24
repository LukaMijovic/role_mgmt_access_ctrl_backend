package model

import "time"

type AccessLog struct {
	logID      int64
	action     string
	accessDate time.Time
	unlockDate time.Time
	userID     int64
	deviceID   int64
}

func (a *AccessLog) GetID() int64 {
	return a.logID
}

func (a *AccessLog) GetAction() string {
	return a.action
}

func (a *AccessLog) GetAccessDate() time.Time {
	return a.accessDate
}

func (a *AccessLog) GetUnlockDate() time.Time {
	return a.unlockDate
}

func (a *AccessLog) GetUserID() int64 {
	return a.userID
}

func (a *AccessLog) GetDeviceID() int64 {
	return a.deviceID
}
