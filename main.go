package main

import (
	"my-go-admin/models"
	"my-go-admin/router"
)

func main() {
	// 初始化 gorm.db
	models.NewGormDB()
	r := router.App()
	r.Run(":8081")
}
