package handlers

import (
	"example.com/m/config"
	"example.com/m/models"
	"github.com/gofiber/fiber/v2"
)

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Change send format"})
	}
	config.DB.Create(&book)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"result": "Book added"})
}

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	config.DB.Find(&books)
	return c.JSON(books)
}

/////ХХ///////////ХХ///////////ХХ///////////ХХ//////
/////ХХ///////////ХХ///////////ХХ///////////ХХ//////

func GetBookById(c *fiber.Ctx) error {

	id := c.Params("id")
	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found"})
	}
	return c.JSON(book)
}

/////ХХ///////////ХХ///////////ХХ///////////ХХ//////
/////ХХ///////////ХХ///////////ХХ///////////ХХ//////

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found"})
	}
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Change send format"})
	}

	config.DB.Save(&book)
	return c.JSON(book)
}

/////ХХ///////////ХХ///////////ХХ///////////ХХ//////
/////ХХ///////////ХХ///////////ХХ///////////ХХ//////

func RemoveBook(c *fiber.Ctx) error {
	id := c.Params("id")

	config.DB.Delete(&models.Book{}, id)
	return c.SendStatus(fiber.StatusNoContent)
}
