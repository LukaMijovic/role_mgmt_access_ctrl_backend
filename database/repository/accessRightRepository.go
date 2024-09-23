package repository

import (
	"database/sql"
	"time"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
)

type AccessRightRepository struct {
	db *sql.DB `binding:required`
}

func NewAccessRightRepository() *AccessRightRepository {
	return &AccessRightRepository{
		db: database.GetDatabaseInstance(),
	}
}

func (arr *AccessRightRepository) GetAccessRightIdFromRole(roleId int64) (int64, error) {
	query := `SELECT access_id FROM public."RoleAccessRight" WHERE role_id = $1`
	row := arr.db.QueryRow(query, roleId)

	var accessId int64
	err := row.Scan(&accessId)

	if err != nil {
		return -1, err
	}

	return accessId, nil
}

func (arr *AccessRightRepository) GetAccessNameFromId(accessId int64) (string, error) {
	query := `SELECT access_name FROM public."Access_right" WHERE access_id = $1`
	row := arr.db.QueryRow(query, accessId)
	//fmt.Printf("Access name: %v\n", accessName)

	var accessName string
	err := row.Scan(&accessName)

	if err != nil {
		return "", err
	}

	return accessName, nil
}

func (arr *AccessRightRepository) SaveTempAccessRight(tempRight *model.TempAccessRight) error {
	query := `INSERT INTO public."Temp_access_right" (access_id, start_access_date, end_access_date, user_id) VALUES ($1, $2, $3, $4)`

	stmt, err := arr.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	stmt.QueryRow(tempRight.GetID(), time.Now(), tempRight.GetEndAccessDate(), tempRight.GetUserID())

	return err
}

func (arr *AccessRightRepository) GetTempAccessRightOfUser(userId int64) ([]int64, error) {
	query := `SELECT access_id FROM public."Temp_access_right" WHERE user_id = $1 AND end_access_date > NOW()`
	rows, err := arr.db.Query(query, userId)

	var tempList []int64

	if err != nil {
		return tempList, err
	}

	defer rows.Close()

	for rows.Next() {
		var tempId int64
		err := rows.Scan(&tempId)

		if err != nil {
			return tempList, err
		}

		tempList = append(tempList, tempId)
	}

	err = rows.Err()

	return tempList, err
}
