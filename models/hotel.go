package models

type Hotel struct {
	Name            string `json:"hotelName"`
	RoomCount       int    `json:"roomCount"`
	OoccupiedRooms  int    `json:"roomOcc"`
	CostPerDay      int    `json:"cost"`
}
