package services

import (
	"errors"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database/repository"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model/dto"
)

func SignInUser(u *dto.UserCredentialsDTO) (*dto.UserCredentialsDTO, error) {
	userRepository := repository.NewUserRepository()

	err := userRepository.GetUserIDFromDataBase(u)

	if err != nil {
		return nil, err
	}

	creds, err := userRepository.GetUserCredentials(u)

	if err != nil {
		return nil, err
	}

	if u.Email == creds.Email && u.Password == creds.Password {
		return creds, nil
	}

	return nil, errors.New("Login failed. Bad credentials.")
}

func RegisterUserToDatabase(u *dto.UserCredentialsDTO) error {
	userRepository := repository.NewUserRepository()

	err := userRepository.GetUserIDFromDataBase(u)

	if err != nil {
		return err
	}

	err = userRepository.SaveUserCredentials(u)

	if err != nil {
		return err
	}

	return nil
}

func SaveUserToDatabase(u *model.User) (*model.User, error) {
	userRepository := repository.NewUserRepository()
	u.SetUserRegistraionDate()
	userID, err := userRepository.Save(u)

	if err != nil {
		return nil, err
	}

	u.SetID(userID)

	return u, nil
}
