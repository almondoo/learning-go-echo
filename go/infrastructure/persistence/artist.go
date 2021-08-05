package persistence

import (
	"learning-go-echo/domain/entity"
	"learning-go-echo/domain/repository"

	"gorm.io/gorm"
)

type ArtistRepository struct {
	Conn *gorm.DB
}

func NewArtistRepository(conn *gorm.DB) repository.ArtistRepository {
	return &ArtistRepository{Conn: conn}
}

func (pr *ArtistRepository) FindByID(id uint) (*entity.Artist, error) {
	artist := &entity.Artist{}
	if err := pr.Conn.First(artist, id).Error; err != nil {
		return nil, err
	}
	return artist, nil
}

func (pr *ArtistRepository) FindByEmail(email string) (*entity.Artist, error) {
	artist := &entity.Artist{}
	if err := pr.Conn.Preload("ArtistCategory").Where("email = ?", email).First(artist).Error; err != nil {
		return nil, err
	}
	return artist, nil
}

func (ar *ArtistRepository) FindByUniqueName(uniqueName string) (*entity.Artist, error) {
	artist := &entity.Artist{}

	if err := ar.Conn.Preload("ArtistCategory").Preload("Products").Preload("Products.ProductCategory").Preload("Products.ProductImages").Preload("Products.ProductFavorites").Where("unique_name = ?", uniqueName).Find(&artist).Error; err != nil {
		return nil, err
	}

	return artist, nil
}

//- 作成
func (pr *ArtistRepository) Create(artist *entity.Artist) (*entity.Artist, error) {
	tx := pr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(artist).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return artist, tx.Commit().Error
}

//- 更新
func (pr *ArtistRepository) Update(artist *entity.Artist) (*entity.Artist, error) {
	tx := pr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}
	if err := pr.Conn.Save(artist).Error; err != nil {
		tx.Rollback()
		return nil, ErrSave
	}

	return artist, tx.Commit().Error
}

//- 削除
func (pr *ArtistRepository) Delete(artist *entity.Artist) error {
	if err := pr.Conn.Delete(&artist).Error; err != nil {
		return ErrDelete
	}

	return nil
}
