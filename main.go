package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"youtube/book"
	"youtube/database"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World")

}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
	app.Get("/", helloWorld)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully")

	errDbMigrate := database.DBConn.AutoMigrate(&book.Book{})
	if errDbMigrate != nil {
		fmt.Print("Migration error")
	}
	fmt.Println("Database Migrated")

}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
