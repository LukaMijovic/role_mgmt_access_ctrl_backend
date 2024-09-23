package repository

import (
	"database/sql"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database"
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
