package repository

import (
	"database/sql"
	"strconv"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model/dto"
)

type UserRepository struct {
	db *sql.DB `binding:"required"`
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.GetDatabaseInstance(),
	}
}

func (ur *UserRepository) GetUserIDFromDataBase(u *dto.UserCredentialsDTO) error {
	query := `SELECT user_id FROM public."User" WHERE email = $1`
	row := ur.db.QueryRow(query, u.Email)

	var res string
	err := row.Scan(&res)

	if err != nil {
		return err
	}

	u.User_ID, err = strconv.ParseInt(res, 10, 64)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserCredentials(u *dto.UserCredentialsDTO) (*dto.UserCredentialsDTO, error) {
	query := `SELECT email, password FROM public."User_credential" WHERE user_id = $1`
	row := ur.db.QueryRow(query, u.User_ID)

	var email string
	var password string

	err := row.Scan(&email, &password)

	if err != nil {
		return nil, err
	}

	return &dto.UserCredentialsDTO{User_ID: u.User_ID, Email: email, Password: password}, nil

}

func (ur *UserRepository) SaveUserCredentials(u *dto.UserCredentialsDTO) error {
	query := `INSERT INTO public."User_credential"(user_id, email, password) VALUES ($1, $2, $3) RETURNING user_id`

	stmt, err := ur.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	var user_id int64
	err = stmt.QueryRow(u.User_ID, u.Email, u.Password).Scan(&user_id)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetRoleIdOfUser(userId int64) (int64, error) {
	query := `SELECT role_id FROM public."User" WHERE user_id = $1`
	row := ur.db.QueryRow(query, userId)

	var roleId int64
	err := row.Scan(&roleId)

	if err != nil {
		return -1, err
	}

	return roleId, nil
}

func (ur *UserRepository) Save(u *model.User) (int64, error) {
	query := `INSERT INTO public."User"(firstname, lastname, email, telephone, birthdate, user_registration_date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id`

	stmt, err := ur.db.Prepare(query)

	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	var userID int64
	err = stmt.QueryRow(u.Firstname, u.Lastname, u.Email, u.Telephone, u.Birthdate, u.GetUserRegistrationDate()).Scan(&userID)

	if err != nil {
		return -1, err
	}

	return userID, nil
}
