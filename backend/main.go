package main

import (
	"net/http"
	"strconv"
	"time"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

var todos []models.Todo

func main() {
	r := gin.Default()

	// GET メソッド
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message" : "Welcome to the API development",
		})
	})

	r.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, todos)
	})

	r.GET("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")

		todoID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID",
			})
			return
		}

		for _, todo := range todos {
			if todo.ID == uint(todoID) {
				c.JSON(http.StatusOK, todo)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found",
		})
	})

	// POST メソッド
	r.POST("/todos", func(c *gin.Context) {
		var newTodo models.Todo

		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		newTodo.ID = uint(len(todos) + 1)
		newTodo.CreatedAt = time.Now()

		todos = append(todos, newTodo)

		c.JSON(http.StatusCreated, newTodo)
	})

	r.Run(":8080")
}