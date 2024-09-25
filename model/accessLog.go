package model

import "time"

type AccessLog struct {
	logID      int64
	action     string
	accessDate time.Time
	unlockDate time.Time
	userId     int64
	deviceId   int64
}

func NewAccessLog(action string, userId int64, deviceId int64) *AccessLog {
	return &AccessLog{
		action:     action,
		accessDate: time.Now(),
		userId:     userId,
		deviceId:   deviceId,
	}
}

func (al *AccessLog) NewEventLog(action string, userId int64, deviceId int64) *AccessLog {
	return &AccessLog{
		action:     action,
		unlockDate: time.Now(),
		userId:     userId,
		deviceId:   deviceId,
	}
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
	return a.userId
}

func (a *AccessLog) GetDeviceID() int64 {
	return a.deviceId
}
