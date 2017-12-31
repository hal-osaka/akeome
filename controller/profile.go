package controller

import (
	"net/http"

	"github.com/alberii/akeome/service"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	twitterID := c.Query("id")
	icon, name, err := service.GetProfile(twitterID)
	if err != nil {
		c.JSON(500, nil)
	}
	c.JSON(http.StatusOK, gin.H{
		"icon":     icon,
		"nickname": name,
	})
}
