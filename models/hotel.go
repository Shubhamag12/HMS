package models

type Hotel struct {
	Name          string `json:"name"`
	RoomCount     int    `json:"room_count"`
	OccupiedRooms int    `json:"occupied_rooms"`
	CostPerDay    int    `json:"cost_per_day"`
}
