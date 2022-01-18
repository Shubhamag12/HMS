package models

type Hotel struct {
	Id              int    `json:"id"`
	Name            string `json:"hotelName"`
	RoomCount       int    `json:"roomCount"`
	UnoccupiedRooms int    `json:"roomOcc"`
	CostPerDay      int    `json:"cost"`
}
