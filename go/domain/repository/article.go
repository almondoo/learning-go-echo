package repository

import "learning-go-echo/domain/entity"

type ArticleRepository interface {
	FindByID(uint, []string) (*entity.Article, error)
	FindByConditions(map[string]interface{}) (*entity.Article, error)
	FindsByConditions(map[string]interface{}) ([]*entity.Article, error)
	GetPageList(int, int, []string) ([]*entity.Article, error)
	Create(*entity.Article) (*entity.Article, error)
	Update(*entity.Article) (*entity.Article, error)
	Delete(*entity.Article) error
}
