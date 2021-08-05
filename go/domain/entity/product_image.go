package entity

import (
	"errors"
	"learning-go-echo/interface/validation"

	"gorm.io/gorm"
)

type ProductImage struct {
	ID        uint     `json:"id"`
	ProductID uint     `json:"productId"`
	Product   *Product `json:"product"`
	Image     string   `json:"image" gorm:"size:255;not null" validate:"required,max=255"`
}

//- 初期値入力
func NewProductImage() ProductImage {
	productImage := ProductImage{}

	return productImage
}

func (pi *ProductImage) TableName() string {
	return "product_images"
}

func (pi *ProductImage) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	err = v.Validate(pi)
	if err != nil {
		return err
	}
	if ok := pi.isExistsProduct(tx); !ok {
		return errors.New("作品が存在しません。")
	}
	return
}

func (pi *ProductImage) isExistsProduct(tx *gorm.DB) bool {
	product := &Product{}
	if err := tx.Where("id = ?", pi.ProductID).First(product).Error; err != nil {
		return false
	}
	return true
}
