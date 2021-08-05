package usecase

import (
	"learning-go-echo/domain/service"
	"learning-go-echo/interface/validation"
)

type RoleUsecase interface {
	Get(requestData *validation.UserLoginRequest) (string, error)
	Set(requestData *validation.UserRegisterRequest) (string, error)
}

type roleUsecase struct {
	roleService service.RoleService
}

func NewRoleUsecase(roleService service.RoleService) RoleUsecase {
	return &roleUsecase{roleService: roleService}
}

func (pu *roleUsecase) Get(requestData *validation.UserLoginRequest) (string, error) {
	return "", nil
}

func (pu *roleUsecase) Set(requestData *validation.UserRegisterRequest) (string, error) {
	return "", nil
}
