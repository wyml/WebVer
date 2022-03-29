package models

import (
	"WebVer/database"

	"gorm.io/gorm"
)

type App struct {
	gorm.Model
	Name      string
	Version   string
	Notice    string
	SecretKey string `gorm:"type:varchar(60);not null"`
}

func GetAppInfo(id int) *gorm.DB {
	return database.DB.Model(new(App)).
		Where("ID = ?", id)
}
