package entity

import (
	"errors"
	"learning-go-echo/interface/validation"
	"time"

	"gorm.io/gorm"
)

type Artist struct {
	ID               uint            `json:"id"`
	ArtistCategoryID uint            `json:"artistCategoryId" gorm:"not null" validate:"required"`
	ArtistCategory   *ArtistCategory `json:"artistCategory"`
	Name             string          `json:"name" gorm:"size:100;not null" validate:"required,max=100"`
	UniqueName       string          `json:"uniqueName" gorm:"size:60;not null;unique" validate:"required,max=60"`
	Email            string          `json:"email" gorm:"size:255;not null;unique" validate:"required,max=255"`
	EmailVerifiedAt  *time.Time      `json:"emailVerifiedAt"`
	Password         string          `json:"-" gorm:"size:255;not null" validate:"required,max=255"`
	Icon             string          `json:"icon" gorm:"size:255;" validate:"max=255"`
	Comment          string          `json:"comment" gorm:"size:200" validate:"max=200"`
	CreatedAt        time.Time       `json:"createdAt" gorm:"not null"`
	UpdatedAt        time.Time       `json:"updatedAt" gorm:"not null"`
	Products         []*Product      `json:"products"`
	Articles         []*Article      `json:"articles"`
}

func (a *Artist) TableName() string {
	return "artists"
}

func (a *Artist) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	err = v.Validate(a)
	if err != nil {
		return err
	}
	if ok := a.isArtistCategoryID(tx); !ok {
		return errors.New("artist_categoryが存在しません。")
	}
	if ok := a.isExistsUniqueName(tx); ok {
		return errors.New("unique_nameが既に存在しています。")
	}
	if ok := a.isExistsEmail(tx); ok {
		return errors.New("emailが既に存在しています。")
	}

	return
}

func (a *Artist) isArtistCategoryID(tx *gorm.DB) bool {
	ArtistCategory := &ArtistCategory{}
	if err := tx.Where("id = ?", a.ArtistCategoryID).First(ArtistCategory).Error; err != nil {
		return false
	}
	return true
}

func (a *Artist) isExistsUniqueName(tx *gorm.DB) bool {
	artist := &Artist{}
	if err := tx.Where("unique_name = ?", a.UniqueName).First(artist).Error; err != nil {
		return false
	}
	return true
}

func (a *Artist) isExistsEmail(tx *gorm.DB) bool {
	artist := &Artist{}
	if err := tx.Where("email = ?", a.Email).First(artist).Error; err != nil {
		return false
	}
	return true
}
