package main

import (
	"go-crud/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/create", handlers.CreateItem)

	log.Fatal(http.ListenAndServe(":5050", nil))
}
