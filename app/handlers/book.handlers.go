package handlers

import (
	"net/http"
	"gorm.io/gorm"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	models "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-EthemCuhadar/app/models"
)

// BOOK HANDLERS

// GetAllBooks method
func GetAllBooks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	books := []models.Book{}
	db.Find(&books)
	respondBookJSON(w, http.StatusOK, books)
}

// CreateBook method
func CreateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	book := models.Book{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		respondBookError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&book).Error; err != nil {
		respondBookError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondBookJSON(w, http.StatusCreated, book)
}

// GetBook method
func GetBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	book := getBookOr404(db, &id, w, r)
	if book == nil {
		return
	}
	respondBookJSON(w, http.StatusOK, book)
}

// DeleteBook method
func DeleteBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	book := getBookOr404(db, &id, w, r)
	if book == nil {
		return
	}
	if err := db.Unscoped().Delete(&book).Error; err != nil {
		respondBookError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondBookJSON(w, http.StatusNoContent, nil)
}

// UpdateBook method
func UpdateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	book := getBookOr404(db, &id, w, r)
	if book == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		respondBookError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&book).Error; err != nil {
		respondBookError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondBookJSON(w, http.StatusOK, book)
}

// GetAuthorOfBook method
func GetAuthorOfBook(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	author := models.Author{}
	vars := mux.Vars(r)
	
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	book := getBookOr404(db, &id, w, r)
	if book == nil {
		return
	}
	if err := db.Preload("Books").Where("id = ?", id).Find(&author).Error; err != nil {
		respondAuthorError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondAuthorJSON(w, http.StatusOK, author)
}

// getBookOr404 gets a book instance if exists, or respond the 404 error otherwise
func getBookOr404(db *gorm.DB, id *uint64, w http.ResponseWriter, r *http.Request) *models.Book {
	book := models.Book{}
	if err := db.First(&book, models.Book{ID: id}).Error; err != nil {
		respondBookError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &book
}