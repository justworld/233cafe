package article

import (
	"233cafe/models"
)

func GetArticles(category int, p int, limit int, s string, orderLike bool, orderRead bool) ([]*models.Article, error) {
	articles, err := models.GetArticleList(category, (p-1)*limit, limit, s, orderLike, orderRead)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func GetArticleInfo(id int) (*models.Article, error) {
	article, err := models.GetArticle(id)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func GetArticleAds(id int) (*models.Ad, error) {
	return nil, nil
}

func GetArticleRecommends(id int) ([]*models.Article, error) {
	return nil, nil

}

func AddArticleRead(id int) {
	models.AddRead(id)
}

func AddArticleLike(id int) {
	models.AddLike(id)
}
