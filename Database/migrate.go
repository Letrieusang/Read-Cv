package main

import (
	"awesomeProject4/Module"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:Iamspectre16@tcp(127.0.0.1:3306)/hoa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(Module.Candidates{})
	log.Println("Migrate Success")
}
