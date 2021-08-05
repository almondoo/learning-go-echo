package seed

import (
	"learning-go-echo/domain/entity"

	"gorm.io/gorm"
)

type productImageSeeder struct {
	conn *gorm.DB
}

func NewProductImageSeeder(db *gorm.DB) *productImageSeeder {
	return &productImageSeeder{conn: db}
}

func (pis *productImageSeeder) Seeder() {
	if ok := pis.Exists(); ok {
		return
	}
	pis.create(entity.ProductImage{
		ProductID: 1,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 1,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 1,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 2,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 2,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 2,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 3,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 3,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 3,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 4,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 4,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 4,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 4,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 4,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 4,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 1,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 5,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 5,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 6,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 6,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 6,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 7,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 7,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 7,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 8,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 8,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 8,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 8,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 8,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 8,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 9,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 9,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 9,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 10,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 10,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 10,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 10,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 10,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 10,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 11,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 11,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 11,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 12,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 12,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 12,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 13,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 13,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 13,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 14,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 14,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 14,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 14,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 14,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 14,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 14,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 15,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 15,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 16,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 16,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 16,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 17,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 17,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 17,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 18,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 18,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 18,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 18,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 18,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 18,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 19,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 19,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 19,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 20,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 20,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 20,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 20,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 20,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 20,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 21,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 21,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 21,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 22,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 22,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 22,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 23,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 23,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 23,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 24,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 24,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 24,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 24,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 24,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 24,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 21,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 25,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 25,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 26,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 26,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 26,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 27,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 27,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 27,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 28,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 28,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 28,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 28,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 28,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 28,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 29,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 29,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 29,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 30,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 30,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 30,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 30,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 30,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 30,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 31,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 31,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 31,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 32,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 32,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 32,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 33,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 33,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 33,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 34,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 34,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 34,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 34,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 34,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 34,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 31,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 35,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 35,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 36,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 36,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 36,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 37,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 37,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 37,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 38,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 38,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 38,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 38,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 38,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 38,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 39,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 39,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 39,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 40,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 40,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 40,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 40,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 40,
		Image:     "https://picsum.photos/500/500",
	})
	pis.create(entity.ProductImage{
		ProductID: 40,
		Image:     "https://picsum.photos/500/500",
	})
}

func (pis *productImageSeeder) create(entity entity.ProductImage) error {
	tx := pis.conn.Begin()
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

func (pis *productImageSeeder) Exists() bool {
	entity := entity.ProductImage{}
	if err := pis.conn.First(&entity).Error; err != nil {
		return false
	}

	return true
}
