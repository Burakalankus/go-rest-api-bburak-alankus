package handlers

import (
	"net/http"

	"github.com/Burakalankus/go-rest-api-bburak-alankus/models"
	"github.com/gin-gonic/gin"
)

func GetAuthors(c *gin.Context) {
	var authors []models.Author
	models.DB.Preload("Books").Find(&authors)
	c.JSON(http.StatusOK, authors)
}

func GetAuthorByID(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	result := models.DB.Preload("Books").First(&author, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	c.JSON(http.StatusOK, author)
}

func CreateAuthor(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&author)
	c.JSON(http.StatusCreated, author)
}

func UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	result := models.DB.First(&author, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Save(&author)
	c.JSON(http.StatusOK, author)
}

func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	result := models.DB.First(&author, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	models.DB.Delete(&author)
	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}

