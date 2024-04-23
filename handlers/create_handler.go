package handlers

import (
	"encoding/json"
	"go-crud/models"
	"net/http"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {

	var newCar models.Car
	json.NewDecoder(r.Body).Decode(&newCar)

	err := newCar.Create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCar)

}
