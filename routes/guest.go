package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"database/sql"

	"github.com/Shubhamag12/HMS/conf"
	"github.com/Shubhamag12/HMS/models"
	"github.com/Shubhamag12/HMS/utils"
	"github.com/gorilla/mux"
)

func CreateGuest(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	var guest models.Guest
	err := json.NewDecoder(r.Body).Decode(&guest)
	if err != nil {
		log.Panicln(err)
	}
	insertRes, insertErr := utils.CreateGuest(conf.DBHandle, &guest)
	if insertErr != nil {
		log.Panicln(insertErr)
	}
	json.NewEncoder(w).Encode(insertRes)
}

func GetGuestByID(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	guestId, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Panicln(err)
	}
	var guest models.Guest
	res :=utils.GetGuestByID(conf.DBHandle, guestId)
	scanErr := res.Scan(&guest)
	if scanErr != nil {
		log.Panicln(scanErr)
	}
	json.NewEncoder(w).Encode(guest)
}

func GetAllGuests(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	var guests []models.Guest
	res, err := utils.GetAllGuests(conf.DBHandle)
	if err != nil {
		log.Panicln(err)
	}
	defer res.Close()
	for res.Next() {
		var guest models.Guest
		if err := res.Scan(&guest); err != nil {
			log.Panicln(err)
		}
		guests = append(guests, guest)
	}

	if err := res.Err(); err != nil {
		log.Panicln(err)
	}
	json.NewEncoder(w).Encode(guests)
}

func UpdateCheckOutDate(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	newCheckoutDate, err := time.Parse("2006-01-02", params["date"])
	if err != nil {
		log.Panicln(err)
	}
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Panicln(err)
	}
	res, UpdateErr := utils.UpdateCheckOutDate(conf.DBHandle, newCheckoutDate, id)
	if UpdateErr != nil {
		log.Panicln(UpdateErr)
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteGuest(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	guestId, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Panicln(err)
	}
	res, deleteErr := utils.DeleteGuest(conf.DBHandle, guestId)
	if deleteErr != nil {
		log.Panicln(deleteErr)
	}
	json.NewEncoder(w).Encode(res)
}