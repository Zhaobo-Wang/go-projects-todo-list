package routes

import (
	"github.com/Zhaobo-Wang/go-projects/controllers"
	"github.com/Zhaobo-Wang/go-projects/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 配置 CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// 公共路由 - 不需要认证
	public := r.Group("/api/v1")
	{
		// 认证路由
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}

	// 受保护路由 - 需要认证
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	{
		// Todo 路由
		protected.GET("/todos", controllers.GetTodos)
		protected.POST("/todos", controllers.CreateTodo)
		protected.GET("/todos/:id", controllers.GetTodo)
		protected.PATCH("/todos/:id", controllers.UpdateTodo)
		protected.PUT("/todos/:id", controllers.UpdateTodo)
		protected.DELETE("/todos/:id", controllers.DeleteTodo)

		// 用户路由
		protected.GET("/user-profile", controllers.GetUser)
	}

	return r
}
