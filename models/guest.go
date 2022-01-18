package models

import "time"

type Guest struct {
	Id           int       `json:"id"`
	Payment      int       `json:"payment"`
	Name         string    `json:"name"`
	CheckInDate  time.Time `json:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date"`
	RoomNumber   int       `json:"room_number"`
}
