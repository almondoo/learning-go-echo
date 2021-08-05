package repository

import "learning-go-echo/domain/entity"

type ArtistRepository interface {
	FindByID(uint) (*entity.Artist, error)
	FindByEmail(string) (*entity.Artist, error)
	FindByUniqueName(string) (*entity.Artist, error)
	Create(*entity.Artist) (*entity.Artist, error)
	Update(*entity.Artist) (*entity.Artist, error)
	Delete(*entity.Artist) error
}
