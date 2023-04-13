package main

import (
	"awesomeProject4/Module"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("Config/")
	viper.ReadInConfig()
	sqlPath := viper.GetString("sql_path")
	log.Println(sqlPath)
	dsn := sqlPath
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(Module.Candidates{})
	log.Println("Migrate Success")
}
