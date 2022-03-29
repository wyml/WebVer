package routers

import (
	"WebVer/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/login", controllers.Login)
	v1 := r.Group("/api")
	{
		v1.GET("/GetAppInfo", controllers.GetAppInfo)
	}
	return r
}
