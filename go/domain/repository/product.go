package repository

import "learning-go-echo/domain/entity"

type ProductRepository interface {
	FindByID(uint) (*entity.Product, error)
	GetPageList(int, int) ([]*entity.Product, error)
	GetListByArtist(uint) ([]*entity.Product, error)
	Create(*entity.Product) (*entity.Product, error)
	Update(*entity.Product) (*entity.Product, error)
	Delete(*entity.Product) error
}
