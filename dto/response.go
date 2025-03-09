package dto

import "github.com/Burakalankus/go-rest-api-bburak-alankus/models"
import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type BookResponse struct {
	Data models.Book `json:"data"`
}

type BooksResponse struct {
	Data []models.Book `json:"data"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type BookRequest struct {
	Title           string `json:"title" binding:"required"`
	AuthorID        uint   `json:"author_id" binding:"required"`
	ISBN            string `json:"isbn" binding:"required"`
	PublicationYear int    `json:"publication_year" binding:"required"`
	Description     string `json:"description"`
}

func HandleError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorResponse{Error: message})
	c.Abort()
}
