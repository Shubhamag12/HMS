package models

import "time"

type Guest struct {
	Id           int
	HotelId      int
	PaymentId    int
	Name         string
	CheckInDate  time.Time
	CheckOutDate time.Time
	RoomNumber   int
}
