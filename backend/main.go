package main

import (
	"log"
	"net/http"
	"todo-app/db"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

var todos []models.Todo

func main() {
	db, err := db.Init()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// GET 一覧
	r.GET("/todos", func(c *gin.Context) {
		var todos []models.Todo
		result := db.Find(&todos)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch todos",
			})
			return
		}

		c.JSON(http.StatusOK, todos)
	})

	r.GET("/todos/:id", func(c *gin.Context) {
		var todo models.Todo
		result := db.First(&todo, c.Param("id"))
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Todo not found",
			})
			return
		}

		c.JSON(http.StatusOK, todo)
	})

	// POST 一覧
	r.POST("/todos", func(c *gin.Context) {
		var newTodo models.Todo

		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		result := db.Create(&newTodo)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create todo",
			})
			return
		}

		c.JSON(http.StatusCreated, newTodo)
	})

	// PUT 一覧
	r.PUT("/todos/:id", func(c *gin.Context) {
		var todo models.Todo
		if err := db.First(&todo, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Todo not found",
			})
			return
		}

		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		
		if err := db.Save(&todo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		
		c.JSON(http.StatusOK, todo)
	})

	// DELETE 一覧
	r.DELETE("/todos/:id", func(c *gin.Context) {
		result := db.Delete(&models.Todo{}, c.Param("id"))
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to delete todo",
			})
			return
		}
		c.Status(http.StatusNoContent)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server startup failed: ", err)
	}
}