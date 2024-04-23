package main

import (
	"go-crud/handlers"
	"go-crud/models"
	"log"
	"net/http"
)

func main() {

	err := models.InitDB("database.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer models.Db.Close()

	http.HandleFunc("/create", handlers.CreateItem)

	log.Fatal(http.ListenAndServe(":5050", nil))

}
