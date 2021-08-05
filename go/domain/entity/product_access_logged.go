package entity

import (
	"learning-go-echo/interface/validation"
	"time"

	"gorm.io/gorm"
)

type ProductAccessLogged struct {
	UserID    uint      `json:"userId"`
	User      User      `json:"user"`
	ProductID uint      `json:"productId" gorm:"not null" validate:"required"`
	Product   Product   `json:"product"`
	Ip        string    `json:"ip" gorm:"size:15;not null" validate:"required,max=15"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
}

//- 初期値入力
func NewProductAccessLogged() ProductAccessLogged {
	productAccessLog := ProductAccessLogged{}

	return productAccessLog
}

func (pal *ProductAccessLogged) TableName() string {
	return "product_access_loggeds"
}

func (pal *ProductAccessLogged) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	err = v.Validate(pal)
	if err != nil {
		return err
	}
	return
}
