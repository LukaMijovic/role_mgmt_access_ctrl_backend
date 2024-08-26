package repository

import (
	"database/sql"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
)

type UserRepository struct {
	db *sql.DB `binding:"required"`
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.GetDatabaseInstance(),
	}
}

func (ur *UserRepository) Save(u *model.User) (int64, error) {
	query := `INSERT INTO public."User"(firstname, lastname, email, telephone, birthdate, user_registration_date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id`

	stmt, err := ur.db.Prepare(query)

	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	var user_id int64
	err = stmt.QueryRow(u.Firstname, u.Lastname, u.Email, u.Telephone, u.Birthdate, u.GetUserRegistrationDate()).Scan(&user_id)

	if err != nil {
		return -1, err
	}

	return user_id, nil
}
