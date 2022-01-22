package models

type Guest struct {
	Id           int    `json:"id,omitempty"`
	Payment      int    `json:"payment,omitempty"`
	Name         string `json:"name,omitempty" validate:"required, min=2, max=64"`
	CheckInDate  string `json:"check_in_date,omitempty" validate:"required, dates"`
	CheckOutDate string `json:"check_out_date,omitempty" validate:"required, dates"`
}
