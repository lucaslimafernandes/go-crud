package main

import (
	"go-crud/handlers"
	"go-crud/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	err := models.InitDB("database.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer models.Db.Close()

	// http.HandleFunc("/create", handlers.CreateItem)
	// http.HandleFunc("/list", handlers.ListAll)
	// http.HandleFunc("/list", handlers.List)

	router.HandleFunc("/create", handlers.CreateItem)

	router.HandleFunc("/list", handlers.ListAll)
	router.HandleFunc("/list/{id}", handlers.List)

	router.HandleFunc("/delete/{id}", handlers.DeleteCar)

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":7000", nil))

}
