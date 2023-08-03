package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"go-fiber-postgres/models"
	"go-fiber-postgres/storage"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Respository struct {
	DB *gorm.DB
}

func (r *Respository) CreateBook(context *fiber.Ctx) error {
	book := Book{}

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err

	}

	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "book has been added"})
	return nil
}

func (r *Respository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Book{}

	err := r.DB.Find(&bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "book FETCHED successfully", "data": bookModels})
	return nil
}

func (r *Respository) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Book{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id is required"})
		return nil
	}

	err := r.DB.Delete(&bookModel, id).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not delete book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "book deleted successfully"})
	return nil
}

func (r *Respository) GetBooksByID(context *fiber.Ctx) error {
	bookModel := &models.Book{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id is required"})
		return nil
	}

	fmt.Println("THE ID IS", id)

	err := r.DB.Where("id = ?", id).First(&bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "book FETCHED successfully", "data": bookModel})
	return nil
}

func (r *Respository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBooksByID)
	api.Get("/get_books", r.GetBooks)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("COULD NOT LOAD THE DATABASE")
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("COULD NOT MIGRATE THE DATABASE")
	}

	r := Respository{
		DB: db,
	}
	//fiber is faster than express
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
