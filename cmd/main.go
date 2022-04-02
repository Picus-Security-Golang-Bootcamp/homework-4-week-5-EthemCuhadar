package main

import (
	config "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-EthemCuhadar/app/config"
	app "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-EthemCuhadar/app"
)

var jsonFilename = "../books.json"

func main(){
	config := config.GetConfig()
	app := app.App{}
	app.Initialize(config)
	app.Run("localhost:8080")
}