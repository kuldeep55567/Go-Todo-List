package routes

import (
	"github.com/gin-gonic/gin"
	"web/controllers"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(r *gin.Engine) {
	// API routes group
	api := r.Group("/v1")
	{
		// User routes
		userRoutes := api.Group("/users")
		{
			userRoutes.GET("/", controllers.GetUser)       // GET /api/users
			userRoutes.POST("/", controllers.CreateUser)    // POST /api/users
		}

		// Todo routes
		todoRoutes := api.Group("/todos")
		{
			todoRoutes.GET("/", controllers.GetTodo)        // GET /api/todos
			todoRoutes.POST("/", controllers.CreateTodo)     // POST /api/todos
		}
	}
}