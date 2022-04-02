package handlers

import (
	"encoding/json"
	"net/http"
)

// respondBookJSON makes the response with books as json format
func respondBookJSON(w http.ResponseWriter, status int, books interface{}) {
	response, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondAuthorJSON makes the response with authors as json format
func respondAuthorJSON(w http.ResponseWriter, status int, authors interface{}) {
	response, err := json.Marshal(authors)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondBookError makes the error response with books as json format
func respondBookError(w http.ResponseWriter, code int, message string) {
	respondBookJSON(w, code, map[string]string{"error": message})
}

// respondBookError makes the error response with authors as json format
func respondAuthorError(w http.ResponseWriter, code int, message string) {
	respondAuthorJSON(w, code, map[string]string{"error": message})
}