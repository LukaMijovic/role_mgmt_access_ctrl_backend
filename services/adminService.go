package services

import (
	"errors"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database/repository"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model/dto"
)

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

	if a.Email == creds.Email && a.Password == creds.Password {
		return creds, nil
	}

	return nil, errors.New("Login failed. Bad credentials.")
}
