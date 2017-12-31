package router

import (
	"github.com/alberii/akeome/controller"
	"github.com/gin-gonic/gin"
)

func apiRouter(api *gin.RouterGroup) {

	api.GET("/@:id", controller.GetMessage)
	api.POST("/@:id", controller.CreateMessage)

	api.GET("/profile", controller.GetProfile)
}
