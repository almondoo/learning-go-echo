package seed

import (
	"learning-go-echo/domain/entity"
	"learning-go-echo/infrastructure/security"

	"gorm.io/gorm"
)

type artistSeeder struct {
	conn *gorm.DB
}

func NewArtistSeeder(db *gorm.DB) *artistSeeder {
	return &artistSeeder{conn: db}
}

func (as *artistSeeder) Seeder() {
	pass, err := security.Hash("testuser")
	if err != nil {
		return
	}

	as.create(entity.Artist{
		ArtistCategoryID: 1,
		Name:             "test1",
		UniqueName:       "fdsagaefda",
		Email:            "user1@example.com",
		Password:         pass,
		Icon:             "https://picsum.photos/200/200",
		Comment:          "コメントコメントコメント\nコメント",
	})
	as.create(entity.Artist{
		ArtistCategoryID: 1,
		Name:             "test2",
		UniqueName:       "gdsagdsagh",
		Email:            "user2@example.com",
		Password:         pass,
		Icon:             "https://picsum.photos/200/200",
		Comment:          "コメントコメントコメント\nコメント",
	})
	as.create(entity.Artist{
		ArtistCategoryID: 3,
		Name:             "test3",
		UniqueName:       "jtrjss4",
		Email:            "user3@example.com",
		Password:         pass,
		Icon:             "https://picsum.photos/200/200",
		Comment:          "コメントコメントコメント\nコメント",
	})
}

func (as *artistSeeder) create(entity entity.Artist) error {
	tx := as.conn.Begin()
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
