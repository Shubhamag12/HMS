package models

type Hotel struct {
	Id              int
	Name            string
	RoomCount       int
	UnoccupiedRooms int
	CostPerDay      int
}
