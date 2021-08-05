package entity

import (
	"errors"
	"learning-go-echo/interface/validation"

	"gorm.io/gorm"
)

type ProductFavorite struct {
	ID         uint    `json:"id"`
	UserID     uint    `json:"userId" gorm:"not null" validate:"required"`
	User       User    `json:"user"`
	ProductID  uint    `json:"productId" gorm:"not null" validate:"required"`
	Product    Product `json:"product"`
	IsFavorite uint8   `json:"isFavorite" gorm:"not null" validate:"min=0,max=1"`
}

//- 初期値入力
func NewProductFavorite() ProductFavorite {
	productFavorite := ProductFavorite{}

	return productFavorite
}

func (pf *ProductFavorite) TableName() string {
	return "product_favorites"
}

func (pf *ProductFavorite) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	err = v.Validate(pf)
	if err != nil {
		return err
	}
	if ok := pf.isExistsUser(tx); !ok {
		return errors.New("ユーザーが存在しません。")
	}
	if ok := pf.isExistsProduct(tx); !ok {
		return errors.New("商品が存在しません。")
	}
	return
}

func (pf *ProductFavorite) isExistsUser(tx *gorm.DB) bool {
	user := &User{}
	if err := tx.Where("id = ?", pf.UserID).First(user).Error; err != nil {
		return false
	}
	return true
}

func (pf *ProductFavorite) isExistsProduct(tx *gorm.DB) bool {
	product := &Product{}
	if err := tx.Where("id = ?", pf.ProductID).First(product).Error; err != nil {
		return false
	}
	return true
}
