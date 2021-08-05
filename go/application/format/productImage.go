package format

import (
	"learning-go-echo/domain/entity"
)

func ProductImagesToArray(productImages []*entity.ProductImage) []string {
	if len(productImages) == 0 {
		return []string{}
	}
	images := make([]string, len(productImages))
	for i, value := range productImages {
		images[i] = value.Image
	}
	return images
}
