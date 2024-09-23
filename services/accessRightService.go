package services

import "github.com/LukaMijovic/role-mgmt-access-ctrl/database/repository"

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

	return accessId == id, nil
}
