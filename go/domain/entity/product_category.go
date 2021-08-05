package entity

import (
	"errors"
	"learning-go-echo/interface/validation"

	"gorm.io/gorm"
)

type ProductCategory struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name" gorm:"size:30;unique;not null" validate:"required,max=30"`
	Products []Product `json:"products"`
}

//- 初期値入力
func NewProductCategory() ProductCategory {
	productCategory := ProductCategory{}

	return productCategory
}

func (pg *ProductCategory) TableName() string {
	return "product_categories"
}

func (pg *ProductCategory) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	err = v.Validate(pg)
	if err != nil {
		return err
	}
	if ok := pg.isExistsName(tx); ok {
		return errors.New("既に同じnameが存在しています。")
	}
	return
}

func (pg *ProductCategory) isExistsName(tx *gorm.DB) bool {
	ArtistCategory := &ArtistCategory{}
	if err := tx.Where("name = ?", pg.Name).First(ArtistCategory).Error; err != nil {
		return false
	}
	return true
}
