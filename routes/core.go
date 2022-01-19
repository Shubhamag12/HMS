package route

import (
	"github.com/Shubhamag12/HMS/utils"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/hotel", utils.GetHotelStats()).Methods("GET")
	router.HandleFunc("/hotel", utils.ChangeTotalRooms()).Methods("PUT")
	router.HandleFunc("/guests", utils.GetAllGuests()).Methods("GET")
	router.HandleFunc("/guests", utils.CreateGuest()).Methods("POST")
	router.HandleFunc("/guests/{id}", utils.GetGuestsByID()).Methods("GET")
	router.HandleFunc("/guests/{id}", utils.UpdateCheckOutDate()).Methods("PUT")
	router.HandleFunc("/guests/{id}", utils.DeleteGuest()).Methods("DELETE")

	return router
}
