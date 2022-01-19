package models

import "time"

type Guest struct {
	Id           int       `json:"id,omitempty"`
	Payment      int       `json:"payment,omitempty"`
	Name         string    `json:"name,omitempty"`
	CheckInDate  time.Time `json:"check_in_date,omitempty"`
	CheckOutDate time.Time `json:"check_out_date,omitempty"`
	RoomNumber   int       `json:"room_number,omitempty"`
}
