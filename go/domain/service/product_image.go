package service

import (
	"learning-go-echo/domain/entity"
	"learning-go-echo/domain/repository"
)

type ProductImageService interface {
	CreateProductImage([]string, uint) ([]*entity.ProductImage, error)
	EditProductImage([]string, uint) ([]*entity.ProductImage, error)
}

type productImageService struct {
	productImageRepo repository.ProductImageRepository
}

func NewProductImageService(productImageRepo repository.ProductImageRepository) ProductImageService {
	return &productImageService{productImageRepo: productImageRepo}
}
func (pis *productImageService) CreateProductImage(images []string, productId uint) ([]*entity.ProductImage, error) {
	productImages := make([]*entity.ProductImage, len(images))

	for index, value := range images {
		productImages[index] = &entity.ProductImage{ProductID: productId, Image: value}
	}

	return pis.productImageRepo.Creates(productImages)
}

func (pis *productImageService) EditProductImage(images []string, productId uint) ([]*entity.ProductImage, error) {
	imageCount := len(images)
	productImages := make([]*entity.ProductImage, imageCount)

	data, err := pis.productImageRepo.FindsByConditions(map[string]interface{}{
		"product_id": productId,
	})
	if err != nil {
		return nil, err
	}

	for index, productImage := range data {
		if index > imageCount {
			pis.productImageRepo.Delete(productImage)
			continue
		}
		productImage.Image = images[index]
		productImage, err := pis.productImageRepo.Update(productImage)
		if err != nil {
			return nil, err
		}
		productImages[index] = productImage
		images = images[1:]
	}

	pdimg, err := pis.CreateProductImage(images, productId)
	if err != nil {
		return nil, err
	}
	productImages = append(productImages, pdimg...)

	return productImages, nil
}
