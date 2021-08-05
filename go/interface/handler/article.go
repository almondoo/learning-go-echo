package handler

import (
	"fmt"
	"learning-go-echo/application/format"
	"learning-go-echo/application/usecase"
	"learning-go-echo/interface/context"
	"learning-go-echo/interface/validation"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleHandler interface {
	GetArticleList() echo.HandlerFunc
	GetArticleArtistList() echo.HandlerFunc
	GetArticleDetail() echo.HandlerFunc
	CreateArticle() echo.HandlerFunc
	EditArticle() echo.HandlerFunc
}

type articleHandler struct {
	au usecase.ArticleUsecase
}

//- ArticleHandlerを生成する
func NewArticleHandler(au usecase.ArticleUsecase) ArticleHandler {
	return &articleHandler{au: au}
}

//- 作品を一覧を取得
func (ah *articleHandler) GetArticleList() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		if page == 0 {
			page = 1
		}
		articles, err := ah.au.GetList(page)
		if err != nil {
			return c.CustomResponse(http.StatusBadRequest, map[string]interface{}{
				"main": "取得できませんでした。",
			})
		}

		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"articles": format.ArticleListFormat(articles),
		})
	})
}

func (ah *articleHandler) GetArticleArtistList() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		page, _ := strconv.Atoi(c.Param("page"))
		articles, err := ah.au.GetList(page)
		if err != nil {
			return c.CustomResponse(http.StatusNotFound, map[string]interface{}{
				"main": "取得できませんでした。",
			})
		}

		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"articles": articles,
		})
	})
}

// - 作品の詳細を取得
func (ah *articleHandler) GetArticleDetail() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		id, _ := strconv.Atoi(c.Param("id"))
		article, err := ah.au.GetDetail(uint(id))
		if err != nil {
			return c.CustomResponse(http.StatusNotFound, map[string]interface{}{
				"main": "取得できませんでした。",
			})
		}

		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"article": format.ArticleDetailFormat(article),
		})
	})
}

//- 作品を作成する
func (ah *articleHandler) CreateArticle() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := new(validation.CreateArticleRequest)
		if ok, message := c.BindValidate(request, validation.CreateArticleMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}

		if err := ah.au.CreateArticle(request, c.GetAuthorID()); err != nil {
			fmt.Println(err)
			return c.CustomResponse(http.StatusInternalServerError, map[string]interface{}{
				"main": "記事の作成に失敗しました。",
			})
		}

		return c.CustomResponse(http.StatusOK, "create")
	})
}

//- 作品を編集する
func (ah *articleHandler) EditArticle() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := new(validation.EditArticleRequest)
		if ok, message := c.BindValidate(request, validation.EditArticleMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}

		if err := ah.au.EditArticle(request, c.GetAuthorID()); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, map[string]interface{}{
				"main": "保存に失敗しました。",
			})
		}
		return c.CustomResponse(http.StatusOK, "edit")
	})
}
