package services

import (
	"errors"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database/repository"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
)

func CheckDeviceIMEIofUser(IMEI string, userId int64) (bool, error) {
	deviceRepository := repository.NewDeviceRepository()
	IMEIList, err := deviceRepository.GetDeviceIMEIOfUser(userId)

	if err != nil {
		return false, err
	}

	if len(IMEIList) == 0 {
		return false, errors.New("There arent any IMEIs for user provided.")
	}

	if len(IMEIList) == 1 {
		return IMEIList[0] == IMEI, nil
	}

	for i := 0; i < len(IMEIList); i++ {
		if IMEIList[i] == IMEI {
			return true, nil
		}
	}

	return false, nil
}

func SaveDeviceToDatabase(d *model.Device) (*model.Device, error) {
	deviceRepository := repository.NewDeviceRepository()
	d.SetDeviceRegistrationDate()
	deviceID, err := deviceRepository.Save(d)

	if err != nil {
		return nil, err
	}

	d.SetID(deviceID)

	return d, nil
}
