package models

import (
	"WebVer/database"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	User   string
	Passwd string
}

type AuthClaims struct {
	ID   int
	User string
	jwt.StandardClaims
}

func GetAdminInfo(user string) *gorm.DB {
	return database.DB.Model(new(Admin)).
		Where("user = ?", user)
}
