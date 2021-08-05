package format

import (
	"learning-go-echo/domain/entity"
)

func ArticleListFormat(articles []*entity.Article) interface{} {
	if len(articles) == 0 {
		return []interface{}{}
	}
	format := make([]interface{}, len(articles))
	for index, article := range articles {

		format[index] = map[string]interface{}{
			"id":        article.ID,
			"artist":    ArtistFormat(article.Artist),
			"title":     article.Title,
			"thumbnail": article.Thumbnail,
			"text":      article.Text,
			"createdAt": article.CreatedAt.Format("2006-1-2"),
		}
	}
	return format
}

func ArticleDetailFormat(article *entity.Article) interface{} {
	return map[string]interface{}{
		"id":        article.ID,
		"artist":    ArtistFormat(article.Artist),
		"title":     article.Title,
		"thumbnail": article.Thumbnail,
		"text":      article.Text,
		"createdAt": article.CreatedAt.Format("2006-1-2"),
		"updatedAt": article.UpdatedAt.Format("2006-1-2"),
	}
}
