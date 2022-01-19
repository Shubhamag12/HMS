package models

type Hotel struct {
	Id            int    `json:"-"`
	Name          string `json:"name"`
	TotalRooms    int    `json:"total_rooms"`
	OccupiedRooms int    `json:"occupied_rooms"`
	CostPerDay    int    `json:"cost_per_day"`
}
