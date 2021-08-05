package entity

import (
	"errors"
	"learning-go-echo/interface/validation"
	"time"

	"gorm.io/gorm"
)

type DeleteUser struct {
	ID               uint           `json:"id" `
	ArtistCategoryID uint           `json:"artistCategoryId" gorm:"not null" validate:"required"`
	ArtistCategory   ArtistCategory `json:"artistCategory" `
	Name             string         `json:"name" gorm:"size:255;not null" validate:"required,max=255"`
	IsArtist         uint8          `json:"isArtist" gorm:"not null" validate:"min=0,max=1"`
	CreatedAt        time.Time      `json:"createdAt" gorm:"not null"`
	UpdatedAt        time.Time      `json:"updatedAt" gorm:"not null"`
}

func NewDeleteUser() DeleteUser {
	deleteUser := DeleteUser{}

	return deleteUser
}

func (du *DeleteUser) TableName() string {
	return "delete_users"
}

func (du *DeleteUser) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	err = v.Validate(du)
	if err != nil {
		return err
	}
	if ok := du.isExistsProductCategory(tx); !ok {
		return errors.New("作品のジャンルが存在しません。")
	}
	return
}

func (du *DeleteUser) isExistsProductCategory(tx *gorm.DB) bool {
	ArtistCategory := &ArtistCategory{}
	if err := tx.Where("id = ?", du.ArtistCategoryID).First(ArtistCategory).Error; err != nil {
		return false
	}
	return true
}
