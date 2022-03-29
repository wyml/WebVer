package main

import (
	"WebVer/database"
	"WebVer/routers"
)

func main() {
	database.InitDB()
	r := routers.InitRouter()
	r.Run()
}
