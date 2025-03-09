package handlers

import (
	"net/http"

	"github.com/Burakalankus/go-rest-api-bburak-alankus/models"
	"github.com/gin-gonic/gin"
)

func GetReviewsForBook(c *gin.Context) {
	bookID := c.Param("id")
	var reviews []models.Review
	models.DB.Where("book_id = ?", bookID).Find(&reviews)
	c.JSON(http.StatusOK, reviews)
}

func CreateReview(c *gin.Context) {
	var review models.Review

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&review)
	c.JSON(http.StatusCreated, review)
}

func UpdateReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review

	result := models.DB.First(&review, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Save(&review)
	c.JSON(http.StatusOK, review)
}

func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review

	result := models.DB.First(&review, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	models.DB.Delete(&review)
	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}

