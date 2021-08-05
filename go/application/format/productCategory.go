package format

import "learning-go-echo/domain/entity"

func ProductCategoryFormat(productCategory *entity.ProductCategory) interface{} {
	if productCategory == nil {
		return nil
	}
	return map[string]interface{}{
		"id":   productCategory.ID,
		"name": productCategory.Name,
	}
}
