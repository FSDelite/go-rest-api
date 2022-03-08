package controllers

import (
	"strconv"

	"github.com/FSDelite/go-rest-api/database"
	"github.com/FSDelite/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {

	id := c.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id format is wrong",
		})
		return
	}
	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "book not found",
		})
		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()
	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func IndexBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book

	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list books: " + err.Error(),
		})
		return
	}
	c.JSON(200, books)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	err = db.Save(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)

}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "id format is wrong",
		})
		return
	}
	db := database.GetDatabase()

	var book models.Book
	err = db.Delete(&book, newId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book",
		})
		return
	}

	c.Status(204)
}
