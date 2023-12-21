package book

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"youtube/database"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	err := db.Find(&books).Error
	errors.Is(err, gorm.ErrRecordNotFound)
	return c.JSON(&books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503)
		errors.Is(err, gorm.ErrRecordNotFound)
	}
	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	result := db.First(&book, id)

	if book.Title == "" {
		c.Status(500).SendString("No Book Found with given id")
		return errors.New("No Book Found with given id")
	}
	errors.Is(result.Error, gorm.ErrRecordNotFound)
	db.Delete(&book)
	return c.SendString("book successfully deleted")
}
