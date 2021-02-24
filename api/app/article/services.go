package article

import (
	"233cafe/models"
)

func GetArticles() (int64, error) {
	count, err := models.GetArticleTotal()
	if err != nil {
		return 0, err
	}

	return count, nil
}