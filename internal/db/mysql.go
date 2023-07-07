package db

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	dsn := "root:lwh123.+@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		log.Fatal(err)
	}
}

func DB(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
