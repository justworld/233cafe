package models

type Config struct {
	Model

	Code string `gorm:"column:code"`
	Data string `gorm:"column:data"`
}

func (v Config) TableName() string {
	return "config"
}
