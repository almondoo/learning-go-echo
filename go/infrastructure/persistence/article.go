package persistence

import (
	"learning-go-echo/domain/entity"
	"learning-go-echo/domain/repository"

	"gorm.io/gorm"
)

type articleRepository struct {
	Conn *gorm.DB
}

func NewArticleRepository(conn *gorm.DB) repository.ArticleRepository {
	return &articleRepository{Conn: conn}
}

//- IDで取得
func (pr *articleRepository) FindByID(id uint, preLoads []string) (*entity.Article, error) {
	article := &entity.Article{ID: id}

	//- Eager Loading
	tx := pr.Conn
	for _, value := range preLoads {
		tx = tx.Preload(value)
	}

	if err := tx.First(&article).Error; err != nil {
		return nil, err
	}

	return article, nil
}

func (ar *articleRepository) FindByConditions(where map[string]interface{}) (*entity.Article, error) {
	var article *entity.Article

	if err := ar.Conn.Where(where).Find(&article).Error; err != nil {
		return nil, err
	}

	return article, nil
}

func (ar *articleRepository) FindsByConditions(where map[string]interface{}) ([]*entity.Article, error) {
	var articles []*entity.Article

	if err := ar.Conn.Where(where).Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (pr *articleRepository) GetPageList(id int, limit int, preLoads []string) ([]*entity.Article, error) {
	var articles []*entity.Article

	//- Eager Loading
	tx := pr.Conn
	for _, value := range preLoads {
		tx = tx.Preload(value)
	}

	if err := tx.Where("id > ?", id).Limit(limit).Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

//- 作成
func (pr *articleRepository) Create(article *entity.Article) (*entity.Article, error) {
	tx := pr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&article).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return article, tx.Commit().Error
}

//- 更新
func (pr *articleRepository) Update(article *entity.Article) (*entity.Article, error) {
	tx := pr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Save(&article).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return article, tx.Commit().Error
}

//- 削除
func (pr *articleRepository) Delete(article *entity.Article) error {
	tx := pr.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(&article).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}
