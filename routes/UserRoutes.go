package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pipoteex/restApiGo/controllers"
)

func UserRoutes(router *gin.Engine){
	routes := router.Group("api/v1/user")
	routes.GET("", controllers.GetUser)
	routes.POST("", controllers.PostUser)
}

