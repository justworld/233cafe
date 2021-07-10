package models

import (
	"233cafe/config"
	"fmt"
	"gorm.io/driver/mysql"
	"log"
	"time"

	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Model struct {
	ID         uint      `gorm:"column:id"`
	Enable     bool      `gorm:"column:enable"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

// Setup initializes the database instance
func Setup() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("root:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DatabaseSetting.Password, config.DatabaseSetting.Host, config.DatabaseSetting.Name), // data source name
		DefaultStringSize:         256,                                                                 // default size for string fields
		DisableDatetimePrecision:  true,                                                                // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                               // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
