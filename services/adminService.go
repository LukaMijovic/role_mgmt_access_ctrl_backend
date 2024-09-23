package services

import (
	"errors"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database/repository"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model/dto"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/util"
)

func ConfirmCreationByAdmin(u *dto.UserCredentialsDTO) error {
	// send signal
	// call addRole...
	return nil
}

func ConfirmTempRightByAdmin(tr *model.TempAccessRight) error {
	//send signal
	return nil
}

func SignInAdmin(a *dto.AdminCredentialsDTO) (*dto.UserCredentialsDTO, error) {
	userRepository := repository.NewUserRepository()

	u := dto.UserCredentialsDTO{
		Email:    a.Email,
		Password: a.Password,
	}

	err := userRepository.GetUserIDFromDataBase(&u)

	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewAdminRepository()
	creds, err := adminRepository.GetAdminCredentials(&u)

	if err != nil {
		return nil, err
	}

	if a.Email == creds.Email && util.CheckPassword(u.Password, creds.Password) {
		return creds, nil
	}

	return nil, errors.New("Login failed. Bad credentials.")
}
