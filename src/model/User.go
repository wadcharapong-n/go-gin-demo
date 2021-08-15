package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID		uint 	`json:"id" gorm:"primary_key"`
	Name 	string 	`json:"name"`
	Age  	int    	`json:"age"`
}