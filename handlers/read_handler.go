package handlers

import (
	"encoding/json"
	"go-crud/models"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {

	// ret := "Retorno"
	var listCars models.ListCars
	json.NewDecoder(r.Body).Decode(&listCars)

	err := listCars.ListAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(ret)
	json.NewEncoder(w).Encode(listCars)

}
