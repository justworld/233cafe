package models

import (
	"233cafe/app"
	"gorm.io/gorm"
)

type Tag struct {
	Model

	Name   string `gorm:"column:name"`
	Online bool   `gorm:"column:online"`
}

type Article struct {
	Model

	Category    int    `gorm:"column:category"`
	Title       string `gorm:"column:title"`
	Desc        string `gorm:"column:desc"`
	HtmlContent string `gorm:"column:html_content"`
	Cover       string `gorm:"column:cover"`
	TagIds      string `gorm:"column:tags"`
	AdIds       string `gorm:"column:ads"`
	IsTop       bool   `gorm:"column:top"`
	ReadNum     int    `gorm:"column:read_num"`
	LikeNum     int    `gorm:"column:like_num"`
	Online      bool   `gorm:"column:online"`
}

type Ad struct {
	Model

	Name   string `gorm:"column:name"`
	Cover  string `gorm:"column:cover"`
	Link   string `gorm:"column:link"`
	Online bool   `gorm:"column:online"`
}

func (v Tag) TableName() string {
	return "tag"
}

func (v Article) TableName() string {
	return "article"
}

func (v Ad) TableName() string {
	return "ad"
}

func GetArticleList(category int, offset int, limit int, s string, orderLike bool, orderRead bool) ([]*Article, error) {
	var articles []*Article
	query := db.Where("enable = ?", 1).Order("create_time desc")
	if category != app.CATEGORY_ALL {
		query = query.Where("category = ?", category)
	}
	if s != "" {
		param := "%" + s + "%"
		query = query.Where("title LIKE ?", param).Or("html_content LIKE ?", param)
	}
	if orderLike {
		query = query.Order("like_num desc")
	}
	if orderRead {
		query = query.Order("read_num desc")
	}
	err := query.Offset(offset).Limit(limit).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ? AND Enable = ? ", id, 1).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

func AddRead(id int) {
	db.Model(&Article{}).Where("id = ?", id).Update("read_num", gorm.Expr(" read_num + ?", 1))
}

func AddLike(id int) {
	db.Model(&Article{}).Where("id = ?", id).Update("like_num", gorm.Expr(" like_num + ?", 1))
}
