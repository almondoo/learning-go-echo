package usecase

import (
	"learning-go-echo/domain/entity"
	"learning-go-echo/domain/service"
	"learning-go-echo/interface/validation"
)

type ArticleUsecase interface {
	GetList(int) ([]*entity.Article, error)
	GetNewList() (string, error)
	GetArtistList() (string, error)
	GetDetail(uint) (*entity.Article, error)
	CreateArticle(*validation.CreateArticleRequest, uint) error
	EditArticle(*validation.EditArticleRequest, uint) error
}

type articleUsecase struct {
	as service.ArticleService
}

func NewArticleUsecase(as service.ArticleService) ArticleUsecase {
	return &articleUsecase{as: as}
}

func (au *articleUsecase) GetList(page int) ([]*entity.Article, error) {
	return au.as.GetArticleList(page)
}

func (au *articleUsecase) GetNewList() (string, error) {
	return "list", nil
}

func (au *articleUsecase) GetArtistList() (string, error) {
	return "list", nil
}

func (au *articleUsecase) GetDetail(id uint) (*entity.Article, error) {
	return au.as.GetArticleDetail(id)
}

func (au *articleUsecase) CreateArticle(request *validation.CreateArticleRequest, artistId uint) error {
	if _, err := au.as.CreateArticle(request, artistId); err != nil {
		return err
	}
	return nil
}

func (au *articleUsecase) EditArticle(request *validation.EditArticleRequest, artistId uint) error {
	_, err := au.as.EditArticle(request, artistId)
	return err
}
