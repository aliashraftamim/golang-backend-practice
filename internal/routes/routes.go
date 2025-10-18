package routes

import (
	"myapp/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/users", handlers.CreateUser)
		// api.GET("/users", handlers.GetUsers)
		// api.GET("/users/:id", handlers.GetUserByID)
		// api.PUT("/users/:id", handlers.UpdateUser)
		// api.DELETE("/users/:id", handlers.DeleteUser)
	}

	return r
}
