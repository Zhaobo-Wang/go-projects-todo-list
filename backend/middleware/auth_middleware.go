package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/Zhaobo-Wang/go-projects/database"
	"github.com/Zhaobo-Wang/go-projects/models"
)

// 从 controllers 包中获取 jwtKey
var jwtKey = []byte(os.Getenv("JWT_SECRET")) 

// Claims 结构体，与 auth_controller.go 中的相同
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// AuthMiddleware 验证 JWT token 并将用户信息添加到上下文
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Authorization 头获取 token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// 检查 Authorization 头格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		// 解析 token
		tokenString := parts[1]
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 查找用户
		var user models.User
		if err := database.DB.First(&user, claims.UserID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// 将用户 ID 存储在上下文中
		c.Set("user_id", claims.UserID)
		c.Set("user", user)

		c.Next()
	}
}

// GetUserID 从上下文中获取当前用户 ID
func GetUserID(c *gin.Context) uint {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0
	}
	return userID.(uint)
}

// GetUser 从上下文中获取当前用户
func GetUser(c *gin.Context) models.User {
	user, exists := c.Get("user")
	if !exists {
		return models.User{}
	}
	return user.(models.User)
}