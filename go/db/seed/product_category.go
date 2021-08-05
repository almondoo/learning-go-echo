package seed

import (
	"learning-go-echo/domain/entity"

	"gorm.io/gorm"
)

type ProductCategorySeeder struct {
	conn *gorm.DB
}

func NewProductCategorySeeder(db *gorm.DB) *ProductCategorySeeder {
	return &ProductCategorySeeder{conn: db}
}

func (pgs *ProductCategorySeeder) Seeder() {
	pgs.create(entity.ProductCategory{
		ID:   1,
		Name: "絵画",
	})
	pgs.create(entity.ProductCategory{
		ID:   2,
		Name: "焼き物",
	})
	pgs.create(entity.ProductCategory{
		ID:   3,
		Name: "写真",
	})
	pgs.create(entity.ProductCategory{
		ID:   4,
		Name: "ドローイング",
	})
	pgs.create(entity.ProductCategory{
		ID:   5,
		Name: "彫刻",
	})
	pgs.create(entity.ProductCategory{
		ID:   6,
		Name: "版画",
	})
}

func (pgs *ProductCategorySeeder) create(entity entity.ProductCategory) error {
	tx := pgs.conn.Begin()
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
