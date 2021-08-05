package entity

import (
	"errors"
	"learning-go-echo/interface/validation"
	"time"

	"gorm.io/gorm"
)

type ArtistAccessLog struct {
	UserID    uint      `json:"userId"`
	User      *User     `json:"user"`
	ArtistID  uint      `json:"artistId" gorm:"not null" validate:"required"`
	Artist    *User     `json:"artist"`
	Ip        string    `json:"ip" gorm:"size:15;not null" validate:"required,max=15"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
}

//- 初期値入力
func NewArtistAccessLog() ArtistAccessLog {
	productAccessLog := ArtistAccessLog{}

	return productAccessLog
}

func (aal *ArtistAccessLog) TableName() string {
	return "artist_access_logs"
}

func (aal *ArtistAccessLog) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	err = v.Validate(aal)
	if err != nil {
		return err
	}
	if ok := aal.isExistsUser(tx); !ok {
		return errors.New("ユーザーが存在しません。")
	}
	if ok := aal.isExistsArtist(tx); !ok {
		return errors.New("アーティストが存在しません。")
	}
	return
}

func (aal *ArtistAccessLog) isExistsUser(tx *gorm.DB) bool {
	user := &User{}
	if err := tx.Where("id = ?", aal.UserID).First(user).Error; err != nil {
		return false
	}
	return true
}

func (aal *ArtistAccessLog) isExistsArtist(tx *gorm.DB) bool {
	artist := &User{}
	if err := tx.Where("id = ?", aal.ArtistID).First(artist).Error; err != nil {
		return false
	}
	return true
}
