package model

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDataBase() {
  database, err := gorm.Open("mysql", "admin:admin1234@(menkungdb.cvil833slmkr.ap-southeast-1.rds.amazonaws.com:3306)/go_demo?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&User{})
	// defer database.Close()

  DB = database
}