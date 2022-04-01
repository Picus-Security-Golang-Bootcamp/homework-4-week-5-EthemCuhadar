package app

import (
	"gorm.io/gorm"
	"log"
	database "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-EthemCuhadar/app/database"
	models "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-EthemCuhadar/app/models"
	utils "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-EthemCuhadar/app/utils"
)

type App struct{
	AuthorDB	*gorm.DB
	BookDB		*gorm.DB
}

func (a *App) Initialize(config string){
	db, _ := database.NewPsqlDB(config)
	a.AuthorDB = models.MigrateAuthors(db)
	a.BookDB = models.MigrateBooks(db)
	booklist, err := utils.ReadJSON("../books.json")
	if err != nil{
		log.Println(err)
	}
	database.InsertSampleData(booklist, db)
}