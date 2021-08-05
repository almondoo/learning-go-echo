package entity

import (
	"errors"
	"learning-go-echo/interface/validation"
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID              uint       `json:"id"`
	RoleID          uint       `json:"roleId" gorm:"not null" validate:"required"`
	Role            *Role      `json:"role"`
	Name            string     `json:"name" gorm:"size:255;not null" validate:"required,max=255"`
	Email           string     `json:"email" gorm:"size:255;not null;unique" validate:"required,email,max=255"`
	EmailVerifiedAt *time.Time `json:"emailVerifiedAt"`
	Password        string     `json:"-" gorm:"size:255;not null" validate:"required"`
	CreatedAt       time.Time  `json:"createdAt" gorm:"not null"`
	UpdatedAt       time.Time  `json:"updatedAt" gorm:"not null"`
}

func (a *Admin) TableName() string {
	return "admins"
}

func (a *Admin) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	if err = v.Validate(a); err != nil {
		return err
	}
	if ok := a.isExistsEmail(tx); ok {
		return errors.New("同じメールアドレスが存在しています。")
	}
	return
}

func (a *Admin) isExistsEmail(tx *gorm.DB) bool {
	admin := &Admin{}
	if err := tx.Where("email = ?", a.Email).First(admin).Error; err != nil {
		return false
	}
	return true
}
