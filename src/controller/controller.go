package controller

import (
	"wadcharapong-n/go-gin-demo/src/model"
	"wadcharapong-n/go-gin-demo/src/service"

	"github.com/gin-gonic/gin"
	
)

func GetHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func User(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	u = service.ResUser(u)
	c.JSON(200, u)
}