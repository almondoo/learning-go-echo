package entity

import (
	"errors"
	"learning-go-echo/interface/validation"
	"time"

	"gorm.io/gorm"
)

type ProductAccessLog struct {
	UserID    uint      `json:"userId"`
	User      User      `json:"user"`
	ProductID uint      `json:"productId" gorm:"not null" validate:"required"`
	Product   Product   `json:"product"`
	Ip        string    `json:"ip" gorm:"size:15;not null" validate:"required,max=15"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
}

//- 初期値入力
func NewProductAccessLog() ProductAccessLog {
	productAccessLog := ProductAccessLog{}

	return productAccessLog
}

func (pal *ProductAccessLog) TableName() string {
	return "product_access_logs"
}

func (pal *ProductAccessLog) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	err = v.Validate(pal)
	if err != nil {
		return err
	}
	if ok := pal.isExistsUser(tx); !ok {
		return errors.New("ユーザーが存在しません。")
	}
	if ok := pal.isExistsProduct(tx); !ok {
		return errors.New("商品が存在しません。")
	}
	return
}

func (pal *ProductAccessLog) isExistsUser(tx *gorm.DB) bool {
	user := &User{}
	if err := tx.Where("id = ?", pal.UserID).First(user).Error; err != nil {
		return false
	}
	return true
}

func (pal *ProductAccessLog) isExistsProduct(tx *gorm.DB) bool {
	product := &Product{}
	if err := tx.Where("id = ?", pal.ProductID).First(product).Error; err != nil {
		return false
	}
	return true
}
