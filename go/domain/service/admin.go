package service

import (
	"learning-go-echo/domain/entity"
	"learning-go-echo/domain/repository"
	"learning-go-echo/infrastructure/auth"
	"learning-go-echo/infrastructure/security"
	"learning-go-echo/interface/validation"
)

type AdminService interface {
	Login(request *validation.AdminLoginRequest) (*entity.Admin, error)
	Register(request *validation.AdminRegisterRequest) (*entity.Admin, error)
	Logout(*auth.TokenData) error
}

type adminService struct {
	adminRepo repository.AdminRepository
	auth      auth.AuthInterface
	token     auth.TokenInterface
}

func NewAdminService(adminRepo repository.AdminRepository, auth auth.AuthInterface, token auth.TokenInterface) AdminService {
	return &adminService{adminRepo: adminRepo, auth: auth, token: token}
}

func (as *adminService) Login(request *validation.AdminLoginRequest) (*entity.Admin, error) {
	admin, err := as.adminRepo.FindByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if err := security.VerifyPassword(admin.Password, request.Password); err != nil {
		return nil, err
	}

	return admin, nil
}

func (as *adminService) Register(request *validation.AdminRegisterRequest) (*entity.Admin, error) {
	password, err := security.Hash(request.Password)
	if err != nil {
		return nil, err
	}
	admin := &entity.Admin{
		RoleID:   20,
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
	}
	admin, err = as.adminRepo.CreateAdmin(admin)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (as *adminService) Logout(tokens *auth.TokenData) error {
	detailtoken, err := as.token.ExtractTokenMetadata(tokens.AccessToken)
	if err != nil {
		return err
	}
	if err := as.auth.DeleteTokens(detailtoken); err != nil {
		return err
	}
	return nil
}
