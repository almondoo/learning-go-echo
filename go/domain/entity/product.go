package entity

import (
	"errors"
	"learning-go-echo/interface/validation"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID                uint               `json:"id"`
	ArtistID          uint               `json:"artistId" gorm:"not null" validate:"required"`
	Artist            *Artist            `json:"artist"`
	ProductCategoryID uint               `json:"productCategoryId" gorm:"not null" validate:"required"`
	ProductCategory   *ProductCategory   `json:"productCategory"`
	Title             string             `json:"title" gorm:"size:255;not null" validate:"required,max=255"`
	Thumbnail         string             `json:"thumbnail" gorm:"size:255;not null" validate:"required,max=255"`
	Price             uint               `json:"price" gorm:"not null" validate:"required,min=1,max=1000000000000000000"`
	Description       string             `json:"description" gorm:"size:text;not null" validate:"required,max=3000"`
	VerticalLength    uint32             `json:"verticalLength" gorm:"not null" validate:"required,max=999999"`
	HorizontalLength  uint32             `json:"horizontalLength" gorm:"not null" validate:"required,max=999999"`
	IsPublished       uint8              `json:"isPublished" gorm:"not null;default:0"`
	IsSold            uint8              `json:"isSold" gorm:"not null;default:0"`
	CreatedAt         time.Time          `json:"createdAt" gorm:"not null"`
	UpdatedAt         time.Time          `json:"updatedAt" gorm:"not null"`
	ProductImages     []*ProductImage    `json:"productImages"`
	ProductFavorites  []*ProductFavorite `json:"productFavorites"`
}

//- 初期値入力
func NewProduct() Product {
	product := Product{}

	return product
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) BeforeSave(tx *gorm.DB) (err error) {
	v := validation.DBValidatorInit()
	err = v.Validate(p)
	if err != nil {
		return err
	}
	if ok := p.isExistsArtist(tx); !ok {
		return errors.New("ユーザーが存在しません。")
	}
	if ok := p.isExistsProductCategory(tx); !ok {
		return errors.New("作品のジャンルが存在しません。")
	}
	return
}

func (p *Product) isExistsArtist(tx *gorm.DB) bool {
	user := &User{}
	if err := tx.Where("id = ?", p.ArtistID).First(user).Error; err != nil {
		return false
	}
	return true
}

func (p *Product) isExistsProductCategory(tx *gorm.DB) bool {
	ProductCategory := &ProductCategory{}
	if err := tx.Where("id = ?", p.ArtistID).First(ProductCategory).Error; err != nil {
		return false
	}
	return true
}
