package seed

import (
	"learning-go-echo/domain/entity"

	"gorm.io/gorm"
)

type articleSeeder struct {
	conn *gorm.DB
}

func NewArticleSeeder(db *gorm.DB) *articleSeeder {
	return &articleSeeder{conn: db}
}

func (as *articleSeeder) Seeder() {
	as.create(entity.Article{
		ID:        1,
		ArtistID:  1,
		Title:     "タイトル",
		Thumbnail: "https://picsum.photos/1000/660",
		Text:      "<div><p>テキスト</p><p>テキスト</p><p>テキスト</p><p>テキスト</p><p>テキスト</p><p>テキスト</p></div>",
	})
	as.create(entity.Article{
		ID:        2,
		ArtistID:  1,
		Title:     "タイトル2",
		Thumbnail: "https://picsum.photos/1000/660",
		Text:      "<div><p>テキスト</p><p>テキスト</p><p>テキスト</p><p>テキスト</p><p>テキスト</p><p>テキスト</p></div>",
	})
	as.create(entity.Article{
		ID:        3,
		ArtistID:  1,
		Title:     "タイトル3",
		Thumbnail: "https://picsum.photos/1000/660",
		Text:      "<div><p>テキスト</p><p>テキスト</p><p>テキスト</p><p>テキスト</p><p>テキスト</p><p>テキスト</p></div>",
	})

}

func (as *articleSeeder) create(entity entity.Article) error {
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
