package model

import "time"

type Device struct {
	deviceID              int64
	IMEI                  string `binding:"required"`
	deviceRegistraionDate time.Time
	UserID                int64 `binding:"required"`
}

func (d *Device) GetID() int64 {
	return d.deviceID
}

func (d *Device) SetID(deviceID int64) {
	d.deviceID = deviceID
}

func (d *Device) GetDeviceRegistraionDate() time.Time {
	return d.deviceRegistraionDate
}

func (d *Device) SetDeviceRegistrationDate() {
	d.deviceRegistraionDate = time.Now()
}
