package models

type Guest struct {
	Id           int    `json:"id,omitempty"`
	Payment      int    `json:"payment,omitempty"`
	Name         string `json:"name,omitempty"`
	CheckInDate  string `json:"check_in_date,omitempty"`
	CheckOutDate string `json:"check_out_date,omitempty"`
}


func (g *Guest) IsEmpty() bool {
	return g.Name == "" || g.CheckInDate == "" || g.CheckOutDate == ""
}