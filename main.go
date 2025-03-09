package main

import (
	"log"
	"github.com/Burakalankus/go-rest-api-bburak-alankus/handlers"
	"github.com/Burakalankus/go-rest-api-bburak-alankus/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect db
	models.ConnectDatabase()

	// Migrations start!
	models.Migrate()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Go REST API is running!",
		})
	})

	// Books endpoints
	bookRoutes := router.Group("/api/v1/books")
	{
		bookRoutes.GET("", handlers.GetBooks)
		bookRoutes.GET("/:id", handlers.GetBookByID)
		bookRoutes.POST("", handlers.CreateBook)
		bookRoutes.PUT("/:id", handlers.UpdateBook)
		bookRoutes.DELETE("/:id", handlers.DeleteBook)
	}

	authorRoutes := router.Group("/api/v1/authors")
	{
		authorRoutes.GET("", handlers.GetAuthors)
		authorRoutes.GET("/:id", handlers.GetAuthorByID)
		authorRoutes.POST("", handlers.CreateAuthor)
		authorRoutes.PUT("/:id", handlers.UpdateAuthor)
		authorRoutes.DELETE("/:id", handlers.DeleteAuthor)
	}
	
	log.Println("Server is running on port 8080")
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Server failed to start")
	}
}

