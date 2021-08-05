package model

import "learning-go-echo/domain/entity"

type ProductListWithArtist struct {
	Artist   *entity.Artist
	Products []*entity.Product
}
