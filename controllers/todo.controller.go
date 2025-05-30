package controllers

import (
	"context"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"web/configs"
	"web/models"
)

// GetTodo fetches a single todo from the database
func GetTodo(c *gin.Context) {
	// Get database connection
	db := configs.GetDB()
	ctx := context.Background()

	// Query to get a single todo
	var todo models.Todo
	err := db.QueryRow(ctx, "SELECT id, user_id, title, description, completed, created_at, updated_at FROM todos LIMIT 1").Scan(
		&todo.ID,
		&todo.UserID,
		&todo.Title,
		&todo.Description,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error fetching todo: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to fetch todo",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   todo,
	})
}

// CreateTodo creates a new todo in the database
func CreateTodo(c *gin.Context) {
	// Parse request body
	var request models.CreateTodoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	// Get database connection
	db := configs.GetDB()
	ctx := context.Background()

	// Insert the new todo
	var todoID int
	err := db.QueryRow(ctx,
		"INSERT INTO todos (user_id, title, description, completed) VALUES ($1, $2, $3, false) RETURNING id",
		request.UserID, request.Title, request.Description,
	).Scan(&todoID)

	if err != nil {
		log.Printf("Error creating todo: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create todo",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Todo created successfully",
		"id":      todoID,
	})
}