package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Zhaobo-Wang/go-projects/database"
	"github.com/Zhaobo-Wang/go-projects/models"
)

// GetTodos fetches all todos
/**
func GetTodos(c *gin.Context)：定义一个函数，接收 Gin 的上下文参数
var todos []models.Todo：声明一个 Todo 结构体的切片（类似数组）
database.DB.Find(&todos)：使用 GORM 查询数据库中所有 Todo 记录，存入 todos 变量
&todos 是传递 todos 的指针，这样函数可以修改它的值
c.JSON(http.StatusOK, gin.H{"data": todos})：
返回 HTTP 200 状态码
gin.H{...} 创建一个 map，键为 "data"，值为 todos 切片
将这个 map 转换为 JSON 并发送给客户端
**/
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	database.DB.Find(&todos)
	
	c.JSON(http.StatusOK, gin.H{"data": todos})
}

// CreateTodo creates a new todo
/**
var input models.Todo：声明一个 Todo 结构体变量
c.ShouldBindJSON(&input)：
尝试将请求体中的 JSON 绑定到 input 变量
如果失败（例如格式错误），返回错误
todo := models.Todo{...}：创建一个新的 Todo 对象，使用从请求中获取的数据
database.DB.Create(&todo)：将新 Todo 保存到数据库
c.JSON(http.StatusCreated, gin.H{"data": todo})：返回 201 状态码和新创建的 Todo
**/
func CreateTodo(c *gin.Context) {
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	todo := models.Todo{
		Title:       input.Title,
		Description: input.Description,
		Completed:   input.Completed,
	}
	
	database.DB.Create(&todo)
	c.JSON(http.StatusCreated, gin.H{"data": todo})
}

// GetTodo fetches a single todo
/**
var todo models.Todo：声明一个 Todo 变量
database.DB.Where("id = ?", c.Param("id")).First(&todo).Error：
c.Param("id")：获取 URL 路径中的 id 参数
Where("id = ?", ...)：构建 SQL WHERE 条件
First(&todo)：查找第一个匹配的记录
.Error：获取可能的错误
如果有错误（找不到记录），返回 404 状态码
否则，返回找到的 Todo 对象
**/
func GetTodo(c *gin.Context) {
	var todo models.Todo
	
	if err := database.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// UpdateTodo updates a todo
/**
首先查找要更新的 Todo，如果不存在返回 404
然后从请求体绑定 JSON 数据到 input 变量
database.DB.Model(&todo).Updates(...)：
Model(&todo) 指定要更新的模型
Updates(...) 使用新数据更新记录
返回更新后的 Todo
**/
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	
	if err := database.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	database.DB.Model(&todo).Updates(models.Todo{
		Title:       input.Title,
		Description: input.Description,
		Completed:   input.Completed,
	})
	
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// DeleteTodo deletes a todo
/**
首先查找要删除的 Todo，如果不存在返回 404
database.DB.Delete(&todo)：从数据库中删除这个 Todo
返回成功消息
**/
func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	
	if err := database.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	
	database.DB.Delete(&todo)
	
	c.JSON(http.StatusOK, gin.H{"data": "Todo deleted successfully"})
}

/**
Go 语言中的 err != nil 解释
err != nil 是 Go 语言中最常见的错误处理模式，用于检查一个操作是否产生了错误。

基本概念
nil：Go 语言中的"空值"，类似于其他语言中的 null、None 或 undefined
!=：不等于操作符
err：通常是一个 error 类型的变量，用于接收函数返回的错误信息
工作原理
在 Go 中，很多函数会返回两个值：

操作的结果
一个表示错误的 error 类型值
如果操作成功，错误值为 nil；如果操作失败，错误值会包含错误信息。
**/