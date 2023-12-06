package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pipoteex/restApiGo/config"
	"github.com/pipoteex/restApiGo/models"
)

type UserType struct{
	Name string `json:"name" binding:"required"`
	Age int `json:"age" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func GetUser(ctx *gin.Context) {

	var user []models.User
	config.DB.Find(&user)
	ctx.JSON(200, &user)
	return

}

func PostUser(c *gin.Context) {

	body := UserType{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{Name: body.Name, Age: body.Age, Email: body.Email}

	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "error al hacer el post"})
		return
	}

	c.JSON(200, &user)
}



