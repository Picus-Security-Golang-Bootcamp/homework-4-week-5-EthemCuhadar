package app

import (
	"gorm.io/gorm"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	database "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-EthemCuhadar/app/database"
	models "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-EthemCuhadar/app/models"
	utils "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-EthemCuhadar/app/utils"
	handlers "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-EthemCuhadar/app/handlers"
)

// App has router and db instances
type App struct{
	Router		*mux.Router
	AuthorDB	*gorm.DB
	BookDB		*gorm.DB
}

// Initialize with predefined configuration and databases. Also sample
// data will be inserted to database which is created.
func (a *App) Initialize(config string){
	db, _ := database.NewPsqlDB(config)
	a.AuthorDB = models.MigrateAuthors(db)
	a.BookDB = models.MigrateBooks(db)
	booklist, err := utils.ReadJSON("../books.json")
	if err != nil{
		log.Println(err)
	}
	database.InsertSampleData(booklist, db)
	a.Router = mux.NewRouter()
	a.SetRouters()
}

// Set all required routers
func (a *App) SetRouters() {
	// Routing for handling the projects
	a.Get("/books", a.GetAllBooks)
	a.Post("/books", a.CreateBook)
	a.Get("/books/{id}", a.GetBook)
	a.Delete("/books/{id}", a.DeleteBook)
	a.Put("/books/{id}", a.UpdateBook)
	a.Get("/authors", a.GetAllAuthors)
	a.Post("/authors", a.CreateAuthor)
	a.Get("/authors/{id}", a.GetAuthor)
	a.Delete("/authors/{id}", a.DeleteAuthor)
	a.Put("/authors/{id}", a.UpdateAuthor)
	a.Get("/authors/{id}/books", a.GetBooksOfAuthor)
	a.Get("/books/{id}/authors", a.GetAuthorOfBook)
}

// Get the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(http.MethodGet)
}

// Post the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(http.MethodPost)
}

// Put the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(http.MethodPut)
}

// Delete the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(http.MethodDelete)
}

// GetAllBooks Handlers to manage Book Data
func (a *App) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	handlers.GetAllBooks(a.BookDB, w, r)
}

// GetBook method
func (a *App) GetBook(w http.ResponseWriter, r *http.Request) {
	handlers.GetBook(a.BookDB, w, r)
}

// CreateBook method
func (a *App) CreateBook(w http.ResponseWriter, r *http.Request){
	handlers.CreateBook(a.BookDB, w, r)
}

// DeleteBook method
func (a *App) DeleteBook(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteBook(a.BookDB, w, r)
}

// UpdateBook method
func (a *App) UpdateBook(w http.ResponseWriter, r *http.Request) {
	handlers.UpdateBook(a.BookDB, w, r)
}

// GetBooksOfAuthor method
func (a *App) GetBooksOfAuthor(w http.ResponseWriter, r *http.Request) {
	handlers.GetBooksOfAuthor(a.BookDB, w, r)
}

// GetAllAuthors Handlers to manage Book Data
func (a *App) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	handlers.GetAllAuthors(a.AuthorDB, w, r)
}

// GetAuthor method
func (a *App) GetAuthor(w http.ResponseWriter, r *http.Request) {
	handlers.GetAuthor(a.AuthorDB, w, r)
}

// CreateAuthor method
func (a *App) CreateAuthor(w http.ResponseWriter, r *http.Request){
	handlers.CreateAuthor(a.AuthorDB, w, r)
}

// DeleteAuthor method
func (a *App) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteAuthor(a.AuthorDB, w, r)
}

// UpdateAuthor method
func (a *App) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	handlers.UpdateAuthor(a.AuthorDB, w, r)
}

// GetAuthorOfBook method
func (a *App) GetAuthorOfBook(w http.ResponseWriter, r *http.Request) {
	handlers.GetAuthorOfBook(a.AuthorDB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string){
	log.Fatal(http.ListenAndServe(host, a.Router))
}