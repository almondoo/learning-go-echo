package entity

import (
	"errors"
	"learning-go-echo/interface/validation"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID        uint      `json:"id"`
	ArtistID  uint      `json:"artistId" gorm:"not null" validate:"required"`
	Artist    *Artist   `json:"artist"`
	Title     string    `json:"title" gorm:"size:255;not null" validate:"required,max=255"`
	Thumbnail string    `json:"thumbnail" gorm:"size:255;not null" validate:"required,max=255"`
	Text      string    `json:"text" gorm:"size:256;not null" validate:"required"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`
}

func (a *Article) TableName() string {
	return "articles"
}

func (a *Article) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	if err = v.Validate(a); err != nil {
		return err
	}

	if ok := a.isExistsArtist(tx); !ok {
		return errors.New("アーティストが存在しません。")
	}
	return
}

func (a *Article) isExistsArtist(tx *gorm.DB) bool {
	user := &User{}
	if err := tx.Where("id = ?", a.ArtistID).First(user).Error; err != nil {
		return false
	}
	return true
}
