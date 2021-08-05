package seed

import (
	"learning-go-echo/domain/entity"

	"gorm.io/gorm"
)

type productSeeder struct {
	conn *gorm.DB
}

func NewProductSeeder(db *gorm.DB) *productSeeder {
	return &productSeeder{conn: db}
}

func (ags *productSeeder) Seeder() {
	ags.create(entity.Product{
		ID:                1,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "タイトル",
		Thumbnail:         "https://picsum.photos/1000/600",
		Price:             50000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    1000,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                2,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト",
		Thumbnail:         "https://picsum.photos/1000/1000",
		Price:             1000000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    1000,
		HorizontalLength:  1000,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                3,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト2",
		Thumbnail:         "https://picsum.photos/300/700",
		Price:             3000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  700,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                4,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト3",
		Thumbnail:         "https://picsum.photos/300/600",
		Price:             20000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            1,
	})
	ags.create(entity.Product{
		ID:                5,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "タイトル",
		Thumbnail:         "https://picsum.photos/1000/600",
		Price:             50000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    1000,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                6,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト",
		Thumbnail:         "https://picsum.photos/1000/1000",
		Price:             1000000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    1000,
		HorizontalLength:  1000,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                7,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト2",
		Thumbnail:         "https://picsum.photos/300/700",
		Price:             3000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  700,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                8,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト3",
		Thumbnail:         "https://picsum.photos/300/600",
		Price:             20000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            1,
	})
	ags.create(entity.Product{
		ID:                9,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト2",
		Thumbnail:         "https://picsum.photos/300/700",
		Price:             3000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  700,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                10,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト3",
		Thumbnail:         "https://picsum.photos/300/600",
		Price:             20000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            1,
	})
	ags.create(entity.Product{
		ID:                11,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "タイトル",
		Thumbnail:         "https://picsum.photos/1000/600",
		Price:             50000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    1000,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                12,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト",
		Thumbnail:         "https://picsum.photos/1000/1000",
		Price:             1000000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    1000,
		HorizontalLength:  1000,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                13,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト2",
		Thumbnail:         "https://picsum.photos/300/700",
		Price:             3000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  700,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                14,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト3",
		Thumbnail:         "https://picsum.photos/300/600",
		Price:             20000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            1,
	})
	ags.create(entity.Product{
		ID:                15,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "タイトル",
		Thumbnail:         "https://picsum.photos/1000/600",
		Price:             50000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    1000,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                16,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト",
		Thumbnail:         "https://picsum.photos/1000/1000",
		Price:             1000000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    1000,
		HorizontalLength:  1000,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                17,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト2",
		Thumbnail:         "https://picsum.photos/300/700",
		Price:             3000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  700,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                18,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト3",
		Thumbnail:         "https://picsum.photos/300/600",
		Price:             20000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            1,
	})
	ags.create(entity.Product{
		ID:                19,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト2",
		Thumbnail:         "https://picsum.photos/300/700",
		Price:             3000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  700,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                20,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト3",
		Thumbnail:         "https://picsum.photos/300/600",
		Price:             20000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            1,
	})
	ags.create(entity.Product{
		ID:                21,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "タイトル",
		Thumbnail:         "https://picsum.photos/1000/600",
		Price:             50000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    1000,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                22,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト",
		Thumbnail:         "https://picsum.photos/1000/1000",
		Price:             1000000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    1000,
		HorizontalLength:  1000,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                23,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト2",
		Thumbnail:         "https://picsum.photos/300/700",
		Price:             3000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  700,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                24,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト3",
		Thumbnail:         "https://picsum.photos/300/600",
		Price:             20000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            1,
	})
	ags.create(entity.Product{
		ID:                25,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "タイトル",
		Thumbnail:         "https://picsum.photos/1000/600",
		Price:             50000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    1000,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                26,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト",
		Thumbnail:         "https://picsum.photos/1000/1000",
		Price:             1000000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    1000,
		HorizontalLength:  1000,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                27,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト2",
		Thumbnail:         "https://picsum.photos/300/700",
		Price:             3000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  700,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                28,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト3",
		Thumbnail:         "https://picsum.photos/300/600",
		Price:             20000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            1,
	})
	ags.create(entity.Product{
		ID:                29,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト2",
		Thumbnail:         "https://picsum.photos/300/700",
		Price:             3000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  700,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                30,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト3",
		Thumbnail:         "https://picsum.photos/300/600",
		Price:             20000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            1,
	})
	ags.create(entity.Product{
		ID:                31,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "タイトル",
		Thumbnail:         "https://picsum.photos/1000/600",
		Price:             50000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    1000,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                32,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト",
		Thumbnail:         "https://picsum.photos/1000/1000",
		Price:             1000000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    1000,
		HorizontalLength:  1000,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                33,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト2",
		Thumbnail:         "https://picsum.photos/300/700",
		Price:             3000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  700,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                34,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト3",
		Thumbnail:         "https://picsum.photos/300/600",
		Price:             20000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            1,
	})
	ags.create(entity.Product{
		ID:                35,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "タイトル",
		Thumbnail:         "https://picsum.photos/1000/600",
		Price:             50000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    1000,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                36,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト",
		Thumbnail:         "https://picsum.photos/1000/1000",
		Price:             1000000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    1000,
		HorizontalLength:  1000,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                37,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト2",
		Thumbnail:         "https://picsum.photos/300/700",
		Price:             3000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  700,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                38,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト3",
		Thumbnail:         "https://picsum.photos/300/600",
		Price:             20000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            1,
	})
	ags.create(entity.Product{
		ID:                39,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト2",
		Thumbnail:         "https://picsum.photos/300/700",
		Price:             3000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  700,
		IsPublished:       1,
		IsSold:            0,
	})
	ags.create(entity.Product{
		ID:                40,
		ArtistID:          1,
		ProductCategoryID: 1,
		Title:             "テスト3",
		Thumbnail:         "https://picsum.photos/300/600",
		Price:             20000,
		Description:       "説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文説明文",
		VerticalLength:    300,
		HorizontalLength:  600,
		IsPublished:       1,
		IsSold:            1,
	})
}

func (ags *productSeeder) create(entity entity.Product) error {
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
