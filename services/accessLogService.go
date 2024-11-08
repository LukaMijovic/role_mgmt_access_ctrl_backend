package services

import (
	"fmt"
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

func LogRequest(userId int64, event string) (int64, error) {
	deviceRepository := repository.NewDeviceRepository()
	deviceId, err := deviceRepository.GetDeviceIdFromUser(userId)

	if err != nil {
		return -1, err
	}

	accessLogRepository := repository.NewAccessLogRepository()
	logId, err := accessLogRepository.Save(model.NewAccessLog(event, userId, deviceId))

	if err != nil {
		return -1, err
	}

	return logId, nil
}

func UpdateUnlockTime(logId int64, unlockTime time.Time, accessId int64) error {
	accessLogRepository := repository.NewAccessLogRepository()
	return accessLogRepository.SaveUnlockTime(logId, unlockTime, fmt.Sprint(accessId))
}
