package repository

import "learning-go-echo/domain/entity"

type AdminRepository interface {
	CreateAdmin(*entity.Admin) (*entity.Admin, error)
	UpdateAdmin(admin *entity.Admin) (*entity.Admin, error)
	DeleteAdmin(admin *entity.Admin) error
	FindByEmail(email string) (*entity.Admin, error)
	FindByID(id uint) (*entity.Admin, error)
}
