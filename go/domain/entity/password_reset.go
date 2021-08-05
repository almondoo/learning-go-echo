package entity

import (
	"learning-go-echo/interface/validation"
	"time"

	"gorm.io/gorm"
)

type PasswordReset struct {
	Email     string    `json:"email" gorm:"size:255;not null" validate:"required,max=255"`
	Token     string    `json:"token" gorm:"size:255;not null" validate:"required,max=255"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
}

//- 初期値入力
func NewPasswordReset() PasswordReset {
	passwordReset := PasswordReset{}

	return passwordReset
}

func (pr *PasswordReset) TableName() string {
	return "password_resets"
}

func (pr *PasswordReset) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	err = v.Validate(pr)
	if err != nil {
		return err
	}
	return
}
