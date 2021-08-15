package service

import "wadcharapong-n/go-gin-demo/src/model"


func ResUser(user model.User) model.User{
	model.DB.Save(&user)
	return user
}