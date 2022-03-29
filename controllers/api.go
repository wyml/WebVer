package controllers

import (
	"WebVer/database"
	"WebVer/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAppInfo(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("Id"))
	data := make([]*models.App, 0)
	err := database.DB.First(&data, Id).Error
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": data,
	})
}
