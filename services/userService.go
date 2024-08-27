package services

import (
	"github.com/LukaMijovic/role-mgmt-access-ctrl/database/repository"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model/dto"
)

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
	user_id, err := userRepository.Save(u)

	if err != nil {
		return nil, err
	}

	u.SetID(user_id)

	return u, nil
}
