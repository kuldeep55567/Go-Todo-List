package controllers

import (
	"context"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"web/configs"
	"web/models"
)

// GetUser fetches a single user from the database
func GetUser(c *gin.Context) {
	// Get database connection
	db := configs.GetDB()
	ctx := context.Background()

	// Query to get a single user
	var user models.User
	err := db.QueryRow(ctx, "SELECT id, username, email, created_at, updated_at FROM users LIMIT 1").Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		log.Printf("Error fetching user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

// CreateUser creates a new user in the database
func CreateUser(c *gin.Context) {
	// Parse request body
	var request models.CreateUserRequest
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

	// Insert the new user
	var userID int
	err := db.QueryRow(ctx,
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id",
		request.Username, request.Email, request.Password,
	).Scan(&userID)

	if err != nil {
		log.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to create user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "User created successfully",
		"id":      userID,
	})
}
