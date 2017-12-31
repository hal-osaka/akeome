package controller

import (
	"net/http"

	"github.com/alberii/akeome/model"
	"github.com/alberii/akeome/service"
	"github.com/gin-gonic/gin"
)

func CreateMessage(c *gin.Context) {
	message := model.Message{}
	c.BindJSON(&message)
	message.TwitterId = c.Param("id")
	if message.TwitterId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "TwitterID is null",
		})
		return
	}
	if message.Body == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Body is null",
		})
		return
	}

	message = service.Message.Store(message)
	c.JSON(http.StatusOK, message)
}

func GetMessage(c *gin.Context) {
	messages := service.Message.FindByTwitter(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
	})
}
