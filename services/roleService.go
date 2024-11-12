package services

import (
	"github.com/LukaMijovic/role-mgmt-access-ctrl/database/repository"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/model"
)

func GetAllRoles() (*[]model.Role, error) {
	roleRepository := repository.NewRoleRepository()
	roles, err := roleRepository.GetAllRoles()

	if err != nil {
		return nil, err
	}

	return roles, nil
}
