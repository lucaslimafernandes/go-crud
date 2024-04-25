package handlers

import (
	"encoding/json"
	"go-crud/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ListAll(w http.ResponseWriter, r *http.Request) {

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

func List(w http.ResponseWriter, r *http.Request) {

	// get params
	params := mux.Vars(r)
	p_id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var car models.Car

	carResult, err := car.List(p_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(carResult)

}
