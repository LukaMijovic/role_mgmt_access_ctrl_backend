package repository

import (
	"database/sql"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model/dto"
)

type AdminRepository struct {
	db *sql.DB `binding:"required"`
}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{
		db: database.GetDatabaseInstance(),
	}
}

func (ar *AdminRepository) GetAdminCredentials(u *dto.UserCredentialsDTO) (*dto.UserCredentialsDTO, error) {
	query := `SELECT email, password FROM public."User_credential" WHERE user_id = $1`
	row := ar.db.QueryRow(query, u.User_ID)

	var email string
	var password string
	err := row.Scan(&email, &password)

	if err != nil {
		return nil, err
	}

	return &dto.UserCredentialsDTO{User_ID: u.User_ID, Email: email, Password: password}, nil
}
