package services

import (
	"github.com/LukaMijovic/role-mgmt-access-ctrl/database/repository"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
)

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
