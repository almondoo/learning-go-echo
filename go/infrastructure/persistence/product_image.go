package persistence

import (
	"learning-go-echo/domain/entity"
	"learning-go-echo/domain/repository"

	"gorm.io/gorm"
)

type productImageRepository struct {
	Conn *gorm.DB
}

func NewProductImageRepository(conn *gorm.DB) repository.ProductImageRepository {
	return &productImageRepository{Conn: conn}
}

func (pir *productImageRepository) SaveOrCreate(productImage *entity.ProductImage, productId uint) (*entity.ProductImage, error) {
	tx := pir.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.FirstOrCreate(productImage).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return productImage, tx.Commit().Error
}

func (pir *productImageRepository) FindByEmail(email string) (entity.ProductImage, error) {
	var productImage entity.ProductImage

	pir.Conn.Where("email = ?", email).First(&productImage)
	return productImage, nil
}

//- IDで取得
func (pir *productImageRepository) FindByID(id uint) (*entity.ProductImage, error) {
	productImage := &entity.ProductImage{ProductID: id}

	if err := pir.Conn.First(&productImage).Error; err != nil {
		return nil, err
	}

	return productImage, nil
}

func (pir *productImageRepository) FindByConditions(where map[string]interface{}) (*entity.ProductImage, error) {
	var productImage *entity.ProductImage

	if err := pir.Conn.Where(where).Find(&productImage).Error; err != nil {
		return nil, err
	}

	return productImage, nil
}

func (pir *productImageRepository) FindsByConditions(where map[string]interface{}) ([]*entity.ProductImage, error) {
	var productImages []*entity.ProductImage

	if err := pir.Conn.Where(where).Find(&productImages).Error; err != nil {
		return nil, err
	}

	return productImages, nil
}

//- 作成
func (pir *productImageRepository) Create(productImage *entity.ProductImage) (*entity.ProductImage, error) {
	tx := pir.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&productImage).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return productImage, tx.Commit().Error
}

func (pir *productImageRepository) Creates(productImages []*entity.ProductImage) ([]*entity.ProductImage, error) {
	tx := pir.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&productImages).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return productImages, tx.Commit().Error
}

//- 更新
func (pir *productImageRepository) Update(productImage *entity.ProductImage) (*entity.ProductImage, error) {
	tx := pir.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Save(&productImage).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return productImage, tx.Commit().Error
}

//- 削除
func (pir *productImageRepository) Delete(productImage *entity.ProductImage) error {
	tx := pir.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(&productImage).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
