package service

import (
	"learning-go-echo/domain/entity"
	"learning-go-echo/domain/repository"
	"learning-go-echo/interface/validation"
	"learning-go-echo/utils/constant"
)

type ProductService interface {
	CreateProduct(*validation.CreateProductRequest, uint) (*entity.Product, error)
	EditProduct(*validation.EditProductRequest) (*entity.Product, error)
	GetProductList(int) ([]*entity.Product, error)
	GetProductListByArtist(string) (*entity.Artist, error)
	GetProductDetail(uint) (*entity.Product, error)
}

type productService struct {
	productRepo repository.ProductRepository
	artistRepo  repository.ArtistRepository
}

func NewProductService(productRepo repository.ProductRepository, artistRepo repository.ArtistRepository) ProductService {
	return &productService{productRepo: productRepo, artistRepo: artistRepo}
}

func (ps *productService) CreateProduct(request *validation.CreateProductRequest, artistId uint) (*entity.Product, error) {
	product := entity.Product{
		ProductCategoryID: request.ProductCategoryID,
		ArtistID:          artistId,
		Title:             request.Title,
		Thumbnail:         request.Thumbnail,
		Price:             request.Price,
		Description:       request.Description,
		VerticalLength:    request.VerticalLength,
		HorizontalLength:  request.HorizontalLength,
		IsPublished:       request.IsPublished,
	}

	createProduct, err := ps.productRepo.Create(&product)

	if err != nil {
		return nil, err
	}

	return createProduct, nil
}

func (ps *productService) EditProduct(request *validation.EditProductRequest) (*entity.Product, error) {
	return nil, nil
}

func (ps *productService) GetProductList(page int) ([]*entity.Product, error) {
	exclusion := (page - 1) * constant.NumberOfProducts

	return ps.productRepo.GetPageList(exclusion, constant.NumberOfProducts)

}

func (ps *productService) GetProductListByArtist(uniqueName string) (*entity.Artist, error) {
	artist, err := ps.artistRepo.FindByUniqueName(uniqueName)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (ps *productService) GetProductDetail(id uint) (*entity.Product, error) {
	return ps.productRepo.FindByID(id)
}
