package usecase

import (
	"learning-go-echo/domain/entity"
	"learning-go-echo/domain/service"
	"learning-go-echo/interface/validation"
)

type ProductUsecase interface {
	GetList(int) ([]*entity.Product, error)
	GetNewList() (string, error)
	GetArtistList(string) (*entity.Artist, error)
	GetDetail(uint) (*entity.Product, error)
	CreateProduct(*validation.CreateProductRequest, uint) error
	Edit(request *validation.EditProductRequest) error
}

type productUsecase struct {
	ps  service.ProductService
	pis service.ProductImageService
}

func NewProductUsecase(ps service.ProductService, pis service.ProductImageService) ProductUsecase {
	return &productUsecase{ps: ps, pis: pis}
}

func (pu *productUsecase) GetList(page int) ([]*entity.Product, error) {
	return pu.ps.GetProductList(page)
}

func (pu *productUsecase) GetNewList() (string, error) {
	return "list", nil
}

func (pu *productUsecase) GetArtistList(uniqueName string) (*entity.Artist, error) {
	return pu.ps.GetProductListByArtist(uniqueName)
}

func (pu *productUsecase) GetDetail(id uint) (*entity.Product, error) {
	return pu.ps.GetProductDetail(id)
}

func (pu *productUsecase) CreateProduct(request *validation.CreateProductRequest, artistId uint) error {
	product, err := pu.ps.CreateProduct(request, artistId)
	if err != nil {
		return err
	}
	if _, err := pu.pis.CreateProductImage(request.ProductImages, product.ID); err != nil {
		return err
	}

	return nil
}

func (pu *productUsecase) Edit(request *validation.EditProductRequest) error {
	product, err := pu.ps.EditProduct(request)
	if err != nil {
		return err
	}
	if _, err := pu.pis.EditProductImage(request.ProductImages, product.ID); err != nil {
		return err
	}

	return nil
}
