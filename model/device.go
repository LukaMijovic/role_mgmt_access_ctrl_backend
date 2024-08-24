package model

import "time"

type Device struct {
	deviceID              int64
	imei                  string
	deviceRegistraionDate time.Time
	userID                int64
}

func (d *Device) GetID() int64 {
	return d.deviceID
}

func (d *Device) GetIMEI() string {
	return d.imei
}

func (d *Device) GetDeviceRegistraionDate() time.Time {
	return d.deviceRegistraionDate
}

func (d *Device) GetUserID() int64 {
	return d.userID
}
