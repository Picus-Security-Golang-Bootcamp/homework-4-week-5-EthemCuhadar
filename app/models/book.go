// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"context"
	"gorm.io/gorm"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Book book
//
// swagger:model book
type Book struct {
	
	gorm.Model
	
	// authorName
	// Example: Cervantes
	AuthorName  *string	`json:"AuthorName"`
	
	// author
	// Example: {"id":1,"name":"Cervantes"}
	Author *Author `gorm:"OnDelete:SET NULL; foreignKey:AuthorName"`
	
	// id
	// Example: 1
	// Required: true
	ID *uint64 `json:"id" gorm:"primaryKey; unique"`

	// isbn
	// Example: 6409887361
	// Required: true
	ISBN *string `json:"isbn" gorm:"unique"`

	// name
	// Example: Don Quiote
	// Required: true
	Name *string `json:"name"`

	// page number
	// Example: 567
	PageNumber *uint64 `json:"pageNumber"`

	// price
	// Example: 19.99
	// Required: true
	Price *float64 `json:"price"`

	// stock code
	// Example: AYT4695
	// Required: true
	StockCode *string `json:"stockCode" gorm:"unique"`

	// stock number
	// Example: 19
	// Required: true
	StockNumber *uint64 `json:"stockNumber"`
}

// BookList is the type in which books will be stored.
type BookList []Book

// TableName returns the name of book schema in PostgreSql database.
func (Book) TableName() string {
	return "books"
}

// MigrateBooks will create and migrate the tables, and then make the some relationships if necessary
func MigrateBooks(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Book{})
	return db
}

// ToString is string representation of relative fields for Book structs.
func (b *Book) ToString() string {
	return fmt.Sprintf("ID: %v, Name: %s, PageNumber: %v, StockNumber: %v, Price: %v, StockCode: %s, ISBN: %s, AuthorName: %s, AuthorID: %v, CreatedAt: %s",
		*b.ID, *b.Name, *b.PageNumber, *b.StockNumber, *b.Price, *b.StockCode, *b.ISBN, *b.Author.Name, *b.Author.ID, b.CreatedAt.Format("2006-01-02 15:04:05"))
}

// ShowBooks prints relative book information of all books in book list
func (bl BookList) ShowBooks() {
	for _, b := range bl {
		fmt.Sprintf("ID: %v, Name: %s, PageNumber: %v, StockNumber: %v, Price: %v, StockCode: %s, ISBN: %s, AuthorName: %s, AuthorID: %v, CreatedAt: %s",
			*b.ID, *b.Name, *b.PageNumber, *b.StockNumber, *b.Price, *b.StockCode, *b.ISBN, *b.Author.Name, *b.Author.ID, b.CreatedAt.Format("2006-01-02 15:04:05"))
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

// Validate validates this book
func (m *Book) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsbn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStockCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStockNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Book) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Book) validateIsbn(formats strfmt.Registry) error {

	if err := validate.Required("isbn", "body", m.ISBN); err != nil {
		return err
	}

	return nil
}

func (m *Book) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Book) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

func (m *Book) validateStockCode(formats strfmt.Registry) error {

	if err := validate.Required("stockCode", "body", m.StockCode); err != nil {
		return err
	}

	return nil
}

func (m *Book) validateStockNumber(formats strfmt.Registry) error {

	if err := validate.Required("stockNumber", "body", m.StockNumber); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this book based on context it is used
func (m *Book) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Book) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Book) UnmarshalBinary(b []byte) error {
	var res Book
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
