package main

import (
	"demo/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Name = "wjh"
	db.Create(user)

	// Read
	fmt.Println(db.First(user, 1)) // find by id

	// Update - 将 models.UserBasic 的 price 更新为 200
	db.Model(user).Update("Password", "1234")

}
