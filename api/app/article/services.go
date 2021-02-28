package article

import (
	"233cafe/models"
)

func GetArticles(p int, limit int, s string) ([]*models.Article, error) {
	articles, err := models.GetArticleList((p-1)*limit, limit, s, false, false)
	if err != nil {
		return nil, err
	}

	return articles, nil
}