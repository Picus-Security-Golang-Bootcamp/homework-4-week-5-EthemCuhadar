package models

import (
	"fmt"
	"gorm.io/gorm"
)

// Author struct to store name and author id
// for book authors.
type Author struct {
	gorm.Model
	Name 	string 	`json:"name"`
	ID   	uint 	`json:"id" gorm:"unique"`
	Books	[]Book
}

// ToString returns a string representation of the author
func (a *Author) ToString() string {
	return fmt.Sprintf("ID: %v Name: %s Books: %v", a.ID, a.Name, a.Books)
}