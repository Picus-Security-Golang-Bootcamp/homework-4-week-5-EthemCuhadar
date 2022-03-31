package utils

import (
	"os"
	models "github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-EthemCuhadar/app/models"
	"encoding/json"
	"io/ioutil"
)

// ReadJSON returns book slice in which objects in json file are read and converted.
// &book is the address of the variable we want to store our parsed data in.
func ReadJSON(filename string) ([]models.Book, error){
	var books []models.Book
	jsonFile, err := os.Open(filename)
	if err != nil{
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err = json.Unmarshal(byteValue, &books); err != nil{
		return nil, err
	}
	return books, nil
}