package model

import "time"

type Device struct {
	DeviceID              int64
	IMEI                  string `binding:"required"`
	deviceRegistraionDate time.Time
	UserID                int64 `binding:"required"`
}

func NewDevice(deviceId int64, IMEI string, deviceRegistraionDate time.Time, userId int64) *Device {
	return &Device{
		DeviceID:              deviceId,
		IMEI:                  IMEI,
		deviceRegistraionDate: deviceRegistraionDate,
		UserID:                userId,
	}
}

func (d *Device) GetID() int64 {
	return d.DeviceID
}

func (d *Device) SetID(deviceID int64) {
	d.DeviceID = deviceID
}

func (d *Device) GetDeviceRegistraionDate() time.Time {
	return d.deviceRegistraionDate
}

func (d *Device) SetDeviceRegistrationDate() {
	d.deviceRegistraionDate = time.Now()
}
