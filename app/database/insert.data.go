package database

import (
	"gorm.io/gorm"
	models "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-EthemCuhadar/app/models"
)

// InsertSampleData takes a book list and and insert the books
// and authors into schemas that are created in database.
func InsertSampleData(list []models.Book, db *gorm.DB) {
	for _, book := range list {
		db.Where(models.Book{ID: book.ID}).
			Attrs(models.Book{ID: book.ID, Name: book.Name}).
			FirstOrCreate(&book)
	}
}