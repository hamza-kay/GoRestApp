package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber"
	"github.com/hamza-kay/restapi/book"
	"github.com/hamza-kay/restapi/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully  opened")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}
func main() {
	app := fiber.New()
	port := os.Getenv("PORT")
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(port)
}
