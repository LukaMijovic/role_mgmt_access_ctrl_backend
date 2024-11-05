package services

import (
	"errors"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database/repository"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model/dto"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/util"
)

func SignInUser(u *dto.UserCredentialsDTO) (*dto.UserCredentialsDTO, error) {
	userRepository := repository.NewUserRepository()

	_, err := userRepository.GetUserIDFromDataBase(u)

	if err != nil {
		return nil, err
	}

	creds, err := userRepository.GetUserCredentials(u)

	if err != nil {
		return nil, err
	}

	if u.Email == creds.Email && util.CheckPassword(u.Password, creds.Password) {
		return creds, nil
	}

	return nil, errors.New("Login failed. Bad credentials.")
}

func RegisterUserToDatabase(u *dto.UserCredentialsDTO) error {
	userRepository := repository.NewUserRepository()

	_, err := userRepository.GetUserIDFromDataBase(u)

	if err != nil {
		return err
	}

	hashedPass, err := util.HashPassword(u.Password)

	if err != nil {
		return err
	}

	u.Password = hashedPass

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

func GetUserFromDataBase(userId int64) (*model.User, error) {
	userRepository := repository.NewUserRepository()
	user, err := userRepository.Read(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetAllUsersFromDataBase() (*[]model.User, error) {
	userRepository := repository.NewUserRepository()
	users, err := userRepository.ReadAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func CheckDoesRegistrationExists(userCrednetials *dto.UserCredentialsDTO) (bool, error) {
	userRepository := repository.NewUserRepository()
	//fmt.Println(userCrednetials.Email)
	ok, err := userRepository.GetUserIDFromDataBase(userCrednetials)

	if err != nil {
		return false, err
	}

	return ok, nil
}
