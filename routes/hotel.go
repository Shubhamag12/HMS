package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Shubhamag12/HMS/conf"
	"github.com/Shubhamag12/HMS/models"
	"github.com/Shubhamag12/HMS/utils"
)

func GetHotelDetails(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	row, err := utils.GetHotelStats(conf.DBHandle)
	if err != nil {
		log.Panicln(err)
	}
	var hotel models.Hotel
	ScanErr := row.Scan(&hotel)
	if err != nil {
		log.Panicln(ScanErr)
	}
	json.NewEncoder(w).Encode(hotel)
}

func CreateHotel(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	var hotel models.Hotel
	err := json.NewDecoder(r.Body).Decode(&hotel)
	if err != nil {
		log.Panicln(err)
	}

	res, createErr := utils.CreateHotel(conf.DBHandle, &hotel)
	if createErr != nil {
		log.Panicln(createErr)
	}

	json.NewEncoder(w).Encode(res)
}