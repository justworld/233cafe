package models

import (
	"gorm.io/driver/mysql"
	"log"
	"time"

	"gorm.io/gorm"
	_ "gorm.io/driver/mysql"
)

var db *gorm.DB

type Model struct {
	ID         uint
	Enable     bool
	CreatedAt  time.Time
	UpdatedAt time.Time
}

// Setup initializes the database instance
func Setup() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/233cafe?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize: 256, // default size for string fields
		DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
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
