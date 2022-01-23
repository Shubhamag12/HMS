package models

type Hotel struct {
	Id            int    `json:"-"`
	Name          string `json:"name,omitempty"`
	TotalRooms    int    `json:"total_rooms,omitempty"`
	OccupiedRooms int    `json:"occupied_rooms"`
	CostPerDay    int    `json:"cost_per_day,omitempty"`
}

func (h *Hotel) IsEmpty() bool {
	return h.CostPerDay == 0 || h.TotalRooms == 0 || h.Name == ""
}
