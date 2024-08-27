package services

import (
	"github.com/LukaMijovic/role-mgmt-access-ctrl/database/repository"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
)

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
