package repository

import "learning-go-echo/domain/entity"

type ProductImageRepository interface {
	SaveOrCreate(*entity.ProductImage, uint) (*entity.ProductImage, error)
	FindByEmail(email string) (entity.ProductImage, error)
	FindByConditions(map[string]interface{}) (*entity.ProductImage, error)
	FindsByConditions(map[string]interface{}) ([]*entity.ProductImage, error)
	Create(product *entity.ProductImage) (*entity.ProductImage, error)
	Creates([]*entity.ProductImage) ([]*entity.ProductImage, error)
	FindByID(id uint) (*entity.ProductImage, error)
	Update(product *entity.ProductImage) (*entity.ProductImage, error)
	Delete(product *entity.ProductImage) error
}
