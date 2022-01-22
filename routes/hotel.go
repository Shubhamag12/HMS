package routes

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Shubhamag12/HMS/conf"
	"github.com/Shubhamag12/HMS/models"
	"github.com/Shubhamag12/HMS/utils"
)

func GetHotelDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	row, err := utils.GetHotelStats(conf.DBHandle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		log.Println(err)
		return
	}

	var id int
	var name string
	var totalRooms int
	var occupiedRooms int
	var costPerDay int
	ScanErr := row.Scan(&id, &name, &totalRooms, &occupiedRooms, &costPerDay)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		log.Println(ScanErr)
		return
	}
	hotel := &models.Hotel{Id: id, Name: name, TotalRooms: totalRooms, OccupiedRooms: occupiedRooms, CostPerDay: costPerDay}
	json.NewEncoder(w).Encode(hotel)
}

func CreateHotel(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	var hotel models.Hotel
	err := json.NewDecoder(r.Body).Decode(&hotel)
	if hotel.IsEmpty() {
		err := errors.New("invalid details")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		log.Panicln(err)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		log.Println(err)
		return
	}
	res, createErr := utils.CreateHotel(conf.DBHandle, &hotel)
	if createErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(createErr.Error())
		log.Println(createErr)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
