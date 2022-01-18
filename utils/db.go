package utils

import (
	"database/sql"
	"log"

	"github.com/Shubhamag12/HMS/models"
)

// CreateHotel This function is ideally called only once to create our fictional hotel we also assume that it has id=1
// (for now)
func CreateHotel(db *sql.DB, hotel *models.Hotel) (sql.Result, error) {
	query := "INSERT INTO hotel_man.hotel (name, room_count, occupied_rooms, cost_per_day) VALUES (?, ?, ?, ?)"
	res, err := db.Exec(query, hotel.Name, hotel.RoomCount, hotel.OccupiedRooms, hotel.CostPerDay)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(res)
	return res, nil
}

// GetHotelStats This functions gives us the vital stats
func GetHotelStats(db *sql.DB) (*sql.Row, error) {
	const id = 1
	query := "SELECT * from hotel_man.hotel WHERE id=1"
	row := db.QueryRow(query)

	return row, nil
}

func ChangeRoomCount(db *sql.DB, newRoomCount int) (sql.Result, error) {
	query := "UPDATE hotel_man.hotel SET room_count=? WHERE id=1"
	res, err := db.Exec(query, newRoomCount)
	if err != nil {
		log.Panicln(err)
	}
	return res, nil
}

func IncrementOccupiedRooms(db *sql.DB) (sql.Result, error) {
	query := "UPDATE hotel_man.hotel SET room_count=room_count+1 WHERE id=1"
	res, err := db.Exec(query)
	if err != nil {
		log.Panicln(err)
	}
	return res, nil
}

func GetAllGuests(db *sql.DB) (*sql.Rows, error) {
	query := "SELECT * from hotel_man.guest"
	rows, err := db.Query(query)
	if err != nil {
		log.Panicln(err)
	}
	return rows, nil
}

func CreateGuest(db *sql.DB, guest *models.Guest) (sql.Result, error) {
	decrementQuery := "UPDATE hotel_man.hotel SET room_count = room_count - 1 where id = 1"
	selectQuery := "SELECT room_count from hotel_man.hotel where id=1"
	incrementQuery := "UPDATE hotel_man.hotel SET occupied_rooms = CASE WHEN room_count >= 0 THEN occupied_rooms + 1 ELSE occupied_rooms where id = 1"
	insertQuery := "INSERT INTO hotel_man.guest (name, check_in_date, check_out_date) VALUES (?, ?, ?)"
	res, decrErr := db.Exec(decrementQuery)
	if decrErr != nil {
		log.Panicln(decrErr)
	}
	log.Println(res)
	row := db.QueryRow(selectQuery)
	var roomCount int

	selErr := row.Scan(&roomCount)

	if selErr != nil {
		log.Panicln(selErr)
	}
}
