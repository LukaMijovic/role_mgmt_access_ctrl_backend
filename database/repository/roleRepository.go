package repository

import (
	"database/sql"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
)

type RoleRepository struct {
	db *sql.DB `binding:"required"`
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{
		db: database.GetDatabaseInstance(),
	}
}

func (rr *RoleRepository) GetAllRoles() (*[]model.Role, error) {
	query := `SELECT * FROM public."Role"`
	rows, err := rr.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var roles []model.Role

	for rows.Next() {
		var roleID int64
		var roleName string
		var roleDescription string

		rows.Scan(&roleID, &roleName, &roleDescription)

		role := model.NewRole(roleID, roleName, roleDescription)

		roles = append(roles, *role)
	}

	return &roles, nil
}
