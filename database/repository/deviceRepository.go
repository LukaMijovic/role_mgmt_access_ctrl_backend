package repository

import (
	"database/sql"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
)

type DeviceRepository struct {
	db *sql.DB `binding:"required"`
}

func NewDeviceRepository() *DeviceRepository {
	return &DeviceRepository{
		db: database.GetDatabaseInstance(),
	}
}

func (dr *DeviceRepository) Save(d *model.Device) (int64, error) {
	query := `INSERT INTO public."Device"("IMEI", device_registration_date, user_id) VALUES ($1, $2, $3) RETURNING device_id`

	stmt, err := dr.db.Prepare(query)

	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	var deviceID int64
	err = stmt.QueryRow(d.IMEI, d.GetDeviceRegistraionDate(), d.UserID).Scan(&deviceID)

	if err != nil {
		return -1, err
	}

	return deviceID, nil
}

func (dr *DeviceRepository) GetDeviceIMEIOfUser(userId int64) ([]string, error) {
	query := `SELECT "IMEI" FROM public."Device" WHERE user_id = $1`
	rows, err := dr.db.Query(query, userId)

	var IMEIList []string

	if err != nil {
		return IMEIList, err
	}

	defer rows.Close()

	for rows.Next() {
		var IMEI string
		err := rows.Scan(&IMEI)

		if err != nil {
			return IMEIList, err
		}

		IMEIList = append(IMEIList, IMEI)
	}

	err = rows.Err()

	return IMEIList, err

}
