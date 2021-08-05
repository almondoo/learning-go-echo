package usecase

import (
	"learning-go-echo/domain/service"
	"learning-go-echo/infrastructure/auth"
	"learning-go-echo/interface/validation"
)

type AdminUsecase interface {
	Login(*validation.AdminLoginRequest) (*auth.TokenDetails, error)
	Register(*validation.AdminRegisterRequest) (*auth.TokenDetails, error)
	Logout(*auth.TokenData) error
}

type adminUsecase struct {
	as    service.AdminService
	auth  auth.AuthInterface
	token auth.TokenInterface
}

func NewAdminUsecase(as service.AdminService, auth auth.AuthInterface, token auth.TokenInterface) AdminUsecase {
	return &adminUsecase{as: as, auth: auth, token: token}
}

func (au *adminUsecase) Login(request *validation.AdminLoginRequest) (*auth.TokenDetails, error) {
	admin, err := au.as.Login(request)
	if err != nil {
		return nil, err
	}

	token, err := au.token.CreateToken(int(admin.ID))
	if err != nil {
		return nil, err
	}

	if err := au.auth.CreateAuth(int(admin.ID), token); err != nil {
		return nil, err
	}

	return token, nil
}

func (au *adminUsecase) Register(request *validation.AdminRegisterRequest) (*auth.TokenDetails, error) {
	admin, err := au.as.Register(request)
	if err != nil {
		return nil, err
	}

	token, err := au.token.CreateToken(int(admin.ID))
	if err != nil {
		return nil, err
	}

	if err := au.auth.CreateAuth(int(admin.ID), token); err != nil {
		return nil, err
	}
	return token, nil
}

func (au *adminUsecase) Logout(tokens *auth.TokenData) error {
	if err := au.as.Logout(tokens); err != nil {
		return err
	}
	return nil
}
