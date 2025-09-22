package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/Zhaobo-Wang/go-projects/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	
	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))
	
	// Routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/todos", controllers.GetTodos)
		v1.POST("/todos", controllers.CreateTodo)
		v1.GET("/todos/:id", controllers.GetTodo)
		v1.PUT("/todos/:id", controllers.UpdateTodo)
		v1.DELETE("/todos/:id", controllers.DeleteTodo)
	}
	
	return r
}