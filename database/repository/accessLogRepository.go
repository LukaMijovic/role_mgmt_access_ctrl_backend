package repository

import (
	"database/sql"
	"time"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
)

type AccessLogRepository struct {
	db *sql.DB `binding:required`
}

func NewAccessLogRepository() *AccessLogRepository {
	return &AccessLogRepository{
		db: database.GetDatabaseInstance(),
	}
}

func (alr *AccessLogRepository) SaveUnlockTime(logId int64, unlockTime time.Time) error {
	query := `UPDATE public."Access_log" SET unlock_date = $1 WHERE log_id = $2`
	stmt, err := alr.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	stmt.QueryRow(unlockTime, logId)

	return nil
}

func (alr *AccessLogRepository) Save(al *model.AccessLog) (int64, error) {
	query := `INSERT INTO public."Access_log"(action, access_date, unlock_date, user_id, device_id) VALUES ($1, $2, $3, $4, $5) RETURNING log_id`

	stmt, err := alr.db.Prepare(query)

	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	var logId int64
	err = stmt.QueryRow(al.GetAction(), al.GetAccessDate(), al.GetUnlockDate(), al.GetUserID(), al.GetDeviceID()).Scan(&logId)

	if err != nil {
		return -1, err
	}

	return logId, nil
}
