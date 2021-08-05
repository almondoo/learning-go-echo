package entity

import (
	"learning-go-echo/interface/validation"
	"time"

	"gorm.io/gorm"
)

type ArtistAccessLogged struct {
	UserID    uint      `json:"userId" `
	User      User      `json:"user" `
	ArtistID  uint      `json:"artistId" gorm:"not null" validate:"required"`
	Artist    User      `json:"artist" `
	Ip        string    `json:"ip" gorm:"size:15;not null" validate:"required,max=15"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
}

//- 初期値入力
func NewArtistAccessLogged() ArtistAccessLogged {
	productAccessLog := ArtistAccessLogged{}

	return productAccessLog
}

func (pal *ArtistAccessLogged) TableName() string {
	return "artist_access_loggeds"
}

func (aal *ArtistAccessLogged) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	err = v.Validate(aal)
	if err != nil {
		return err
	}
	return
}
