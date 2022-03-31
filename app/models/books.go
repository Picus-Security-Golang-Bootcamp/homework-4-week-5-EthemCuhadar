package models

import (
	"fmt"

	"gorm.io/gorm"
)

// BookList is the type in which books will be stored.
type BookList []Book

// Book is struct type to store information about
// books. Author name field is foreign key for
// Author struct.
type Book struct {
	gorm.Model
	ID          uint    `json:"id" gorm:"unique"`
	Name        string  `json:"name"`
	PageNumber  uint    `json:"pageNumber"`
	StockNumber uint    `json:"stockNumber"`
	Price       float64 `json:"price"`
	StockCode   string  `json:"stockCode" gorm:"unique"`
	ISBN        string  `json:"isbn" gorm:"unique"`
	AuthorName  string  `json:"authorName"`
	Author      Author  `gorm:"OnDelete:SET NULL"`
}

// TableName returns the name of book schema in PostgreSql database.
func (Book) TableName() string {
	return "books"
}

// ToString is string representation of relative fields for Book structs.
func (b *Book) ToString() string {
	return fmt.Sprintf("ID: %v, Name: %s, PageNumber: %v, StockNumber: %v, Price: %v, StockCode: %s, ISBN: %s, AuthorName: %s, AuthorID: %v, CreatedAt: %s",
		b.ID, b.Name, b.PageNumber, b.StockNumber, b.Price, b.StockCode, b.ISBN, b.Author.Name, b.Author.ID, b.CreatedAt.Format("2006-01-02 15:04:05"))
}

// ShowBooks prints relative book information of all books in book list
func (bl BookList) ShowBooks() {
	for _, b := range bl {
		fmt.Sprintf("ID: %v, Name: %s, PageNumber: %v, StockNumber: %v, Price: %v, StockCode: %s, ISBN: %s, AuthorName: %s, AuthorID: %v, CreatedAt: %s",
			b.ID, b.Name, b.PageNumber, b.StockNumber, b.Price, b.StockCode, b.ISBN, b.Author.Name, b.Author.ID, b.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}

// BeforeCreate is a callback that gets called before creating
func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println(b.ID, "is created")
	return
}

// AfterDelete is a callback that gets called after deleting
func (b *Book) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println(b.ID, "is deleted")
	return
}
