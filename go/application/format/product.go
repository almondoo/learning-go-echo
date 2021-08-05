package format

import (
	"learning-go-echo/domain/entity"
)

func ProductListFormat(products []*entity.Product) interface{} {
	if len(products) == 0 {
		return []interface{}{}
	}
	format := make([]interface{}, len(products))
	for index, product := range products {
		format[index] = map[string]interface{}{
			"id":        product.ID,
			"artist":    ArtistFormat(product.Artist),
			"category":  ProductCategoryFormat(product.ProductCategory),
			"title":     product.Title,
			"thumbnail": product.Thumbnail,
			"price":     product.Price,
		}
	}
	return format
}

func ProductListWithArtistFormat(artist *entity.Artist) interface{} {
	if artist == nil {
		return nil
	}
	products := make([]interface{}, len(artist.Products))
	for index, product := range artist.Products {
		products[index] = map[string]interface{}{
			"id":        product.ID,
			"category":  ProductCategoryFormat(product.ProductCategory),
			"title":     product.Title,
			"thumbnail": product.Thumbnail,
			"price":     product.Price,
		}
	}

	return map[string]interface{}{
		"artist":   ArtistFormat(artist),
		"products": products,
	}
}

func ProductDetailFormat(product *entity.Product) interface{} {
	if product == nil {
		return nil
	}
	return map[string]interface{}{
		"id":                    product.ID,
		"artist":                ArtistFormat(product.Artist),
		"category":              ProductCategoryFormat(product.ProductCategory),
		"title":                 product.Title,
		"thumbnail":             product.Thumbnail,
		"price":                 product.Price,
		"description":           product.Description,
		"verticalLength":        product.VerticalLength,
		"horizontalLength":      product.HorizontalLength,
		"isPublished":           product.IsPublished,
		"isSold":                product.IsSold,
		"productImages":         ProductImagesToArray(product.ProductImages),
		"productFavoriteNumber": len(product.ProductFavorites),
	}
}
