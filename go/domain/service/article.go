package service

import (
	"learning-go-echo/domain/entity"
	"learning-go-echo/domain/repository"
	"learning-go-echo/interface/validation"
	"learning-go-echo/utils/constant"
)

type ArticleService interface {
	CreateArticle(*validation.CreateArticleRequest, uint) (*entity.Article, error)
	EditArticle(*validation.EditArticleRequest, uint) (*entity.Article, error)
	GetArticleList(int) ([]*entity.Article, error)
	GetArticleDetail(uint) (*entity.Article, error)
}

type articleService struct {
	articleRepo repository.ArticleRepository
}

func NewArticleService(articleRepo repository.ArticleRepository) ArticleService {
	return &articleService{articleRepo: articleRepo}
}

func (ps *articleService) CreateArticle(request *validation.CreateArticleRequest, artistId uint) (*entity.Article, error) {
	article := entity.Article{
		ArtistID:  artistId,
		Title:     request.Title,
		Thumbnail: request.Thumbnail,
		Text:      request.Text,
	}

	return ps.articleRepo.Create(&article)
}

func (ps *articleService) GetArticleList(page int) ([]*entity.Article, error) {
	exclusion := (page - 1) * constant.NumberOfArticles

	return ps.articleRepo.GetPageList(exclusion, constant.NumberOfArticles, []string{"Artist", "Artist.ArtistCategory"})

}

func (ps *articleService) GetArticleDetail(id uint) (*entity.Article, error) {
	return ps.articleRepo.FindByID(id, []string{"Artist", "Artist.ArtistCategory"})
}

func (as *articleService) EditArticle(request *validation.EditArticleRequest, artistId uint) (*entity.Article, error) {
	article, err := as.articleRepo.FindByConditions(map[string]interface{}{
		"id":        request.Id,
		"artist_id": artistId,
	})
	if err != nil {
		return nil, err
	}
	article.Title = request.Title
	article.Thumbnail = request.Thumbnail
	article.Text = request.Text
	return as.articleRepo.Update(article)
}
