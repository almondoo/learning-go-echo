package seed

import (
	"learning-go-echo/domain/entity"

	"gorm.io/gorm"
)

type ArtistCategorySeeder struct {
	conn *gorm.DB
}

func NewArtistCategorySeeder(db *gorm.DB) *ArtistCategorySeeder {
	return &ArtistCategorySeeder{conn: db}
}

func (ags *ArtistCategorySeeder) Seeder() {
	ags.create(entity.ArtistCategory{
		ID:   1,
		Name: "画家",
	})
	ags.create(entity.ArtistCategory{
		ID:   2,
		Name: "美術家",
	})
	ags.create(entity.ArtistCategory{
		ID:   3,
		Name: "彫刻家",
	})
	ags.create(entity.ArtistCategory{
		ID:   4,
		Name: "陶芸家",
	})
}

func (ags *ArtistCategorySeeder) create(entity entity.ArtistCategory) error {
	tx := ags.conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&entity).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
