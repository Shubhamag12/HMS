package routes

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Shubhamag12/HMS/conf"
	"github.com/Shubhamag12/HMS/models"
	"github.com/Shubhamag12/HMS/utils"
	"github.com/gorilla/mux"
)

func CreateGuest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.ContentLength == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var guest models.Guest

	err := json.NewDecoder(r.Body).Decode(&guest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		log.Println(err)
		return
	}

	if guest.IsEmpty() {
		e := errors.New("empty body")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e.Error())
		log.Println(e)
		return
	}

	checkInDate, err := time.Parse("2006-01-02", guest.CheckInDate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		log.Println(err)
		return
	}
	checkOutDate, err := time.Parse("2006-01-02", guest.CheckOutDate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		log.Println(err)
		return
	}
	delta := int(checkOutDate.Sub(checkInDate).Hours() / 24)
	if delta <= 0 {
		e := errors.New("invalid date range")
		json.NewEncoder(w).Encode(err.Error())
		log.Println(e.Error())
		return
	}

	insertRes, insertErr := utils.CreateGuest(conf.DBHandle, &guest)
	if insertErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(insertErr.Error())
		log.Println(insertErr.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertRes)
}

func GetGuestByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	guestId, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		log.Println(err)
		return
	}
	var guest models.Guest
	res := utils.GetGuestByID(conf.DBHandle, guestId)

	scanErr := res.Scan(&guest.Id, &guest.Payment, &guest.Name, &guest.CheckInDate, &guest.CheckOutDate)
	if scanErr != nil {
		if scanErr.Error() == sql.ErrNoRows.Error() {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(scanErr.Error())
		log.Println(scanErr)
		return
	}
	json.NewEncoder(w).Encode(guest)
}

func GetAllGuests(w http.ResponseWriter, r *http.Request) {
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
		if err := res.Scan(&guest.Id, &guest.Payment, &guest.Name, &guest.CheckInDate, &guest.CheckOutDate); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			log.Println(err)
			return
		}
		guests = append(guests, guest)
	}

	if err := res.Err(); err != nil {
		log.Panicln(err)
	}
	json.NewEncoder(w).Encode(guests)
}

func UpdateCheckOutDate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var reqData map[string]string
	parseErr := json.NewDecoder(r.Body).Decode(&reqData)
	if parseErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(parseErr.Error())
		log.Println(parseErr)
		return
	}
	newCheckoutDate, err := time.Parse("2006-01-02", reqData["date"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		log.Println(err)
		return
	}

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		log.Println(err)
		return
	}
	res, UpdateErr := utils.UpdateCheckOutDate(conf.DBHandle, newCheckoutDate, id)
	if UpdateErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(UpdateErr.Error())
		log.Println(UpdateErr.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func DeleteGuest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	guestId, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		log.Println(err)
		return
	}
	res, deleteErr := utils.DeleteGuest(conf.DBHandle, guestId)
	if deleteErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(deleteErr.Error())
		log.Println(deleteErr)
		return
	}
	json.NewEncoder(w).Encode(res)
}
