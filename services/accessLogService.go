package services

import (
	"github.com/LukaMijovic/role-mgmt-access-ctrl/database/repository"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
)

func LogEvent(userId int64, deviceId int64) (int64, error) {
	accessLogRepository := repository.NewAccessLogRepository()
	logId, err := accessLogRepository.Save(model.NewAccessLog("access", userId, deviceId))

	if err != nil {
		return -1, err
	}

	return logId, nil
}
