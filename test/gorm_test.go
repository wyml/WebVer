package test

import (
	"WebVer/models"
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGorm(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:3306)/wlyz?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.App, 0)
	err = db.Where("Id = ?", 1).Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range data {
		fmt.Println(item)
	}
}
