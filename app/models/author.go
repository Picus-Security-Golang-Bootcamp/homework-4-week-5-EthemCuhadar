// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"context"
	"strconv"
	"gorm.io/gorm"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Author author
//
// swagger:model author
type Author struct {
	
	gorm.Model
	
	// books
	// Example: [book1, book2, book3]
	Books []*Book

	// id
	// Example: 1
	// Required: true
	ID *uint64 `json:"id" gorm:"unique"`

	// name
	// Example: Cervantes
	// Required: true
	Name *string `json:"name" gorm:"primaryKey"`
}

// TableName returns the name of book schema in PostgreSql database.
func (Author) TableName() string {
	return "authors"
}

// MigrateBooks will create and migrate the tables, and then make the some relationships if necessary
func MigrateAuthors(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Author{})
	return db
}

// ToString returns a string representation of the author
func (a *Author) ToString() string {
	return fmt.Sprintf("ID: %v Name: %s Books: %v", a.ID, a.Name, a.Books)
}

// Validate validates this author
func (m *Author) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBooks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Author) validateBooks(formats strfmt.Registry) error {
	if swag.IsZero(m.Books) { // not required
		return nil
	}

	for i := 0; i < len(m.Books); i++ {
		if swag.IsZero(m.Books[i]) { // not required
			continue
		}

		if m.Books[i] != nil {
			if err := m.Books[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("books" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("books" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Author) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Author) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this author based on the context it is used
func (m *Author) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBooks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Author) contextValidateBooks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Books); i++ {

		if m.Books[i] != nil {
			if err := m.Books[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("books" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("books" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Author) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Author) UnmarshalBinary(b []byte) error {
	var res Author
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
