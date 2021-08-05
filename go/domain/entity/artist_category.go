package entity

import (
	"errors"
	"learning-go-echo/interface/validation"

	"gorm.io/gorm"
)

type ArtistCategory struct {
	ID      uint      `json:"id"`
	Name    string    `json:"name" gorm:"size:30;not null;unique" validate:"required,max=30"`
	Artists []*Artist `json:"artists"`
}

//- 初期値入力
func NewArtistCategory() ArtistCategory {
	artistCategory := ArtistCategory{}

	return artistCategory
}

func (ag *ArtistCategory) TableName() string {
	return "artist_categories"
}

func (ag *ArtistCategory) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	err = v.Validate(ag)
	if err != nil {
		return err
	}
	if ok := ag.isExistsName(tx); ok {
		return errors.New("既に同じカテゴリーが存在しています。")
	}
	return
}

func (ag *ArtistCategory) isExistsName(tx *gorm.DB) bool {
	artistCategory := &ArtistCategory{}
	if err := tx.Where("name = ?", ag.Name).First(artistCategory).Error; err != nil {
		return false
	}
	return true
}
