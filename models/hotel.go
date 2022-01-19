package models

type Hotel struct {
	Id            int    `json:"-"`
	Name          string `json:"name,omitempty"`
	TotalRooms    int    `json:"total_rooms,omitempty"`
	OccupiedRooms int    `json:"occupied_rooms,omitempty"`
	CostPerDay    int    `json:"cost_per_day,omitempty"`
}
