package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Go REST API is running!",
		})
	})

	log.Println("Server is running on port 8080")
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Server failed to start")
	}
}

