package models

type Hotel struct {
	Id            int    `json:"-"`
	Name          string `json:"name"`
	RoomCount     int    `json:"room_count"`
	OccupiedRooms int    `json:"occupied_rooms"`
	CostPerDay    int    `json:"cost_per_day"`
}
