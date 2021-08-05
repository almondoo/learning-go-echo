package persistence

import (
	"learning-go-echo/domain/entity"
	"learning-go-echo/domain/repository"

	"gorm.io/gorm"
)

type productRepository struct {
	Conn *gorm.DB
}

func NewProductRepository(conn *gorm.DB) repository.ProductRepository {
	return &productRepository{Conn: conn}
}

//- IDで取得
func (pr *productRepository) FindByID(id uint) (*entity.Product, error) {
	product := &entity.Product{ID: id}

	if err := pr.Conn.Preload("Artist").Preload("Artist.ArtistCategory").Preload("ProductCategory").Preload("ProductImages").Preload("ProductFavorites").First(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (pr *productRepository) GetPageList(id int, limit int) ([]*entity.Product, error) {
	var products []*entity.Product

	if err := pr.Conn.Preload("Artist").Preload("Artist.ArtistCategory").Preload("ProductCategory").Preload("ProductImages").Preload("ProductFavorites").Where("id > ?", id).Limit(limit).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (pr *productRepository) GetListByArtist(artistId uint) ([]*entity.Product, error) {
	var products []*entity.Product

	//- Eager Loading
	if err := pr.Conn.Preload("ProductCategory").Preload("ProductImages").Preload("ProductFavorites").Where("artist_id = ?", artistId).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

//- 作成
func (pr *productRepository) Create(product *entity.Product) (*entity.Product, error) {
	tx := pr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return product, tx.Commit().Error
}

//- 更新
func (pr *productRepository) Update(product *entity.Product) (*entity.Product, error) {
	tx := pr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Save(&product).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return product, tx.Commit().Error
}

//- 削除
func (pr *productRepository) Delete(product *entity.Product) error {
	tx := pr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(&product).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}
