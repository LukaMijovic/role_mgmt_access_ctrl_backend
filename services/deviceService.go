package services

import (
	"errors"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database/repository"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
)

func CheckDeviceIMEIofUser(IMEI string, userId int64) (int64, bool, error) {
	deviceRepository := repository.NewDeviceRepository()
	IMEIList, err := deviceRepository.GetDeviceIMEIOfUser(userId)

	if err != nil {
		return -1, false, err
	}

	if len(IMEIList) == 0 {
		return -1, false, errors.New("There arent any IMEIs for user provided.")
	}

	if len(IMEIList) == 1 {
		deviceId, err := deviceRepository.GetDeviceIdFromIMEI(IMEI)

		if err != nil {
			return -1, IMEIList[0] == IMEI, err
		}

		return deviceId, IMEIList[0] == IMEI, nil
	}

	for i := 0; i < len(IMEIList); i++ {
		if IMEIList[i] == IMEI {
			deviceId, err := deviceRepository.GetDeviceIdFromIMEI(IMEI)

			if err != nil {
				return -1, true, err
			}

			return deviceId, true, nil
		}
	}

	return -1, false, nil
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
