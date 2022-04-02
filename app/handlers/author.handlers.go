package handlers

import (
	"net/http"
	"gorm.io/gorm"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	models "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-EthemCuhadar/app/models"
)

// AUTHOR HANDLERS

// GetAllAuthors method
func GetAllAuthors(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var authors []models.Author
	db.Find(&authors)
	respondAuthorJSON(w, http.StatusOK, authors)
}

// CreateAuthor method
func CreateAuthor(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	author := models.Author{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&author); err != nil {
		respondAuthorError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&author).Error; err != nil {
		respondAuthorError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondAuthorJSON(w, http.StatusCreated, author)
}

// GetAuthor method
func GetAuthor(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	author := getAuthorOr404(db, &id, w, r)
	if author == nil {
		return
	}
	respondAuthorJSON(w, http.StatusOK, author)
}

// DeleteAuthor method
func DeleteAuthor(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	author := getAuthorOr404(db, &id, w, r)
	if author == nil {
		return
	}
	if err := db.Unscoped().Delete(&author).Error; err != nil {
		respondAuthorError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondAuthorJSON(w, http.StatusNoContent, nil)
}

// UpdateAuthor method
func UpdateAuthor(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	author := getAuthorOr404(db, &id, w, r)
	if author == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&author); err != nil {
		respondAuthorError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&author).Error; err != nil {
		respondAuthorError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondAuthorJSON(w, http.StatusOK, author)
}

// GetBooksOfAuthor method
func GetBooksOfAuthor(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	books := []models.Book{}
	vars := mux.Vars(r)
	
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	author := getAuthorOr404(db, &id, w, r)
	if author == nil {
		return
	}
	if err := db.Preload("Author").Where(&models.Book{ID: &id}).Find(&books).Error; err != nil {
		respondBookError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondBookJSON(w, http.StatusOK, books)
}

// getAuthorOr404 gets a author instance if exists, or respond the 404 error otherwise
func getAuthorOr404(db *gorm.DB, id *uint64, w http.ResponseWriter, r *http.Request) *models.Author {
	author := models.Author{}
	if err := db.First(&author, models.Author{ID: id}).Error; err != nil {
		respondAuthorError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &author
}