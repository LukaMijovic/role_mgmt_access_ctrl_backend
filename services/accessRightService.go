package services

import (
	"errors"

	"github.com/LukaMijovic/role-mgmt-access-ctrl/database/repository"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
)

func CheckAccessRightOfUser(userId int64, id int64) (bool, error) {
	userRepository := repository.NewUserRepository()
	roleId, err := userRepository.GetRoleIdOfUser(userId)

	if err != nil {
		return false, err
	}

	accessRightRepository := repository.NewAccessRightRepository()
	accessId, err := accessRightRepository.GetAccessRightIdFromRole(roleId)

	if err != nil {
		return false, err
	}

	if accessId != id {
		tempList, err := accessRightRepository.GetTempAccessRightOfUser(userId)

		if err != nil {
			return false, err
		}

		if len(tempList) == 0 {
			return false, errors.New("No rights found.")
		}

		if len(tempList) == 1 {
			return tempList[0] == id, nil
		}

		for i := 0; i < len(tempList); i++ {
			if tempList[i] == id {
				return true, nil
			}
		}

	}

	return false, nil
}

func GiveTempAccessRightToUser(tempAccess *model.TempAccessRight) error {
	accessRightRepository := repository.NewAccessRightRepository()
	accessName, err := accessRightRepository.GetAccessNameFromId(tempAccess.GetID())

	if err != nil {
		return err
	}

	tempAccess.SetAccessName(accessName)

	err = accessRightRepository.SaveTempAccessRight(tempAccess)

	return err
}
