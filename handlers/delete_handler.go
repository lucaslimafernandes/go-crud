package handlers

import (
	"encoding/json"
	"go-crud/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteCar(w http.ResponseWriter, r *http.Request) {

	var car models.Car
	// get params
	params := mux.Vars(r)
	p_id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = car.Delete(p_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(p_id)

}
