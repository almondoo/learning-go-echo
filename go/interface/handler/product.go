package handler

import (
	"learning-go-echo/application/format"
	"learning-go-echo/application/usecase"
	"learning-go-echo/interface/context"
	"learning-go-echo/interface/validation"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler interface {
	GetProductList() echo.HandlerFunc
	GetProductArtistList() echo.HandlerFunc
	GetProductDetail() echo.HandlerFunc
	CreateProduct() echo.HandlerFunc
	Edit() echo.HandlerFunc
}

type productHandler struct {
	pu usecase.ProductUsecase
}

//- ProductHandlerを生成する
func NewProductHandler(pu usecase.ProductUsecase) ProductHandler {
	return &productHandler{pu: pu}
}

//- 作品を一覧を取得
func (ph *productHandler) GetProductList() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		if page == 0 {
			page = 1
		}
		products, err := ph.pu.GetList(page)
		if err != nil {
			return c.CustomResponse(http.StatusBadRequest, map[string]interface{}{
				"main": "取得できませんでした。",
			})
		}

		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"products": format.ProductListFormat(products),
		})
	})
}

func (ph *productHandler) GetProductArtistList() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		uniqueName := c.Param("artist")
		artist, err := ph.pu.GetArtistList(uniqueName)
		if err != nil {
			return c.CustomResponse(http.StatusNotFound, map[string]interface{}{
				"main": "取得できませんでした。",
			})
		}

		return c.CustomResponse(http.StatusOK, format.ProductListWithArtistFormat(artist))
	})
}

// - 作品の詳細を取得
func (ph *productHandler) GetProductDetail() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		id, _ := strconv.Atoi(c.Param("id"))
		product, err := ph.pu.GetDetail(uint(id))
		if err != nil {
			return c.CustomResponse(http.StatusNotFound, map[string]interface{}{
				"main": "取得できませんでした。",
			})
		}

		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"product": format.ProductDetailFormat(product),
		})
	})
}

//- 作品を作成する
func (ph *productHandler) CreateProduct() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := new(validation.CreateProductRequest)
		if ok, message := c.BindValidate(request, validation.CreateProductMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}

		if err := ph.pu.CreateProduct(request, c.GetAuthorID()); err != nil {
			log.Fatal(err)
			return c.CustomResponse(http.StatusInternalServerError, map[string]interface{}{
				"main": "サーバーエラー",
			})
		}

		return c.CustomResponse(http.StatusOK, "create")
	})
}

//- 作品を編集する
func (ph *productHandler) Edit() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := new(validation.EditProductRequest)
		if ok, message := c.BindValidate(request, validation.EditProductMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}
		if err := ph.pu.Edit(request); err != nil {
			return c.CustomResponse(http.StatusBadRequest, map[string]interface{}{
				"main": "正常に保存されませんでした。",
			})
		}
		return c.CustomResponse(http.StatusOK, "edit")
	})
}

func (ph *productHandler) Delete() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		return c.CustomResponse(http.StatusOK, "delete")
	})
}
