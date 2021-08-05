package format

import "learning-go-echo/domain/entity"

func ArtistFormat(artist *entity.Artist) interface{} {
	if artist == nil {
		return nil
	}
	return map[string]interface{}{
		"id":         artist.ID,
		"category":   ArtistCategoryFormat(artist.ArtistCategory),
		"name":       artist.Name,
		"uniqueName": artist.UniqueName,
		"icon":       artist.Icon,
	}
}
