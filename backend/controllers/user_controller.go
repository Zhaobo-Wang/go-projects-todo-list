package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Zhaobo-Wang/go-projects/middleware"
)

// GetUser 获取当前登录用户的信息
func GetUser(c *gin.Context) {
	user := middleware.GetUser(c)
	
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"createdAt": user.CreatedAt,
			"updatedAt": user.UpdatedAt,
		},
	})
}