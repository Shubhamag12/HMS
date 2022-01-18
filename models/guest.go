package models

import time

type Guest struct {
	Id           int    `json:"id"`
	Payment      int    `json:"payment"`
	Name         string `json:"name"`
	CheckInDate  string `json:"inDate"`
	CheckOutDate string `json:"outDate"`
	RoomNumber   int    `json:"RoomNumber"`
}
