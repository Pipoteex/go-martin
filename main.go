package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pipoteex/restApiGo/config"
	"github.com/pipoteex/restApiGo/models"
	"github.com/pipoteex/restApiGo/routes"
)

func main() {
	r := gin.Default()

	config.ConnectToDB()

	config.DB.AutoMigrate(models.Task{})
	config.DB.AutoMigrate(models.User{})

	routes.UserRoutes(r)
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}