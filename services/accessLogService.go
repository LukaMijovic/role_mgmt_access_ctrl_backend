package services

import (
	"time"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database/repository"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
)

func LogEvent(userId int64, deviceId int64, event string) (int64, error) {
	accessLogRepository := repository.NewAccessLogRepository()
	logId, err := accessLogRepository.Save(model.NewAccessLog(event, userId, deviceId))

	if err != nil {
		return -1, err
	}

	return logId, nil
}

func UpdateUnlockTime(logId int64, unlockTime time.Time) error {
	accessLogRepository := repository.NewAccessLogRepository()
	return accessLogRepository.SaveUnlockTime(logId, unlockTime)
}
