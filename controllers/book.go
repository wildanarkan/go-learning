package controllers

import (
    "wildanarkan.com/go-learning/models"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type BookController struct {
    DB *gorm.DB
}

// Get all books
func (b *BookController) GetBooks(c *gin.Context) {
    var books []models.Book
    b.DB.Find(&books)
    c.JSON(http.StatusOK, gin.H{"data": books})
}

// Get book by id
func (b *BookController) GetBook(c *gin.Context) {
    var book models.Book
    
    if err := b.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": book})
}

// Create new book
func (b *BookController) CreateBook(c *gin.Context) {
    var input models.Book
    
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    book := models.Book{
        Title:       input.Title,
        Author:      input.Author,
        Publisher:   input.Publisher,
        PublishYear: input.PublishYear,
    }
    
    b.DB.Create(&book)
    c.JSON(http.StatusCreated, gin.H{"data": book})
}

// Update book
func (b *BookController) UpdateBook(c *gin.Context) {
    var book models.Book
    
    if err := b.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
        return
    }

    var input models.Book
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    b.DB.Model(&book).Updates(input)
    c.JSON(http.StatusOK, gin.H{"data": book})
}

// Delete book
func (b *BookController) DeleteBook(c *gin.Context) {
    var book models.Book
    
    if err := b.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
        return
    }

    b.DB.Delete(&book)
    c.JSON(http.StatusOK, gin.H{"data": "Book deleted successfully"})
}