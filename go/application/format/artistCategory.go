package format

import "learning-go-echo/domain/entity"

func ArtistCategoryFormat(artistCategory *entity.ArtistCategory) interface{} {
	if artistCategory == nil {
		return nil
	}
	return map[string]interface{}{
		"id":   artistCategory.ID,
		"name": artistCategory.Name,
	}
}
