package models

type Article struct {
	Model

	Title         string
	Desc          string
}

// GetArticleTotal gets the total number of articles based on the constraints
func GetArticleTotal() (int64, error) {
	var count int64
	if err := db.Model(&Article{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}