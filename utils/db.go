package utils

import (
	"database/sql"
	"log"
	"time"

	"github.com/Shubhamag12/HMS/models"
)

// CreateHotel This function is ideally called only once to create our fictional hotel we also assume that it has id=1
// (for now)
func CreateHotel(db *sql.DB, hotel *models.Hotel) (sql.Result, error) {
	query := "INSERT INTO hotel_man.hotel (name, total_rooms, cost_per_day) VALUES (?, ?, ?)"
	res, err := db.Exec(query, hotel.Name, hotel.TotalRooms, hotel.CostPerDay)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(res)
	return res, nil
}

// GetHotelStats This functions gives us the vital stats
func GetHotelStats(db *sql.DB) (*sql.Row, error) {
	query := "SELECT * from hotel_man.hotel WHERE id=1"
	row := db.QueryRow(query)

	return row, nil
}

// ChangeTotalRooms TODO: requires CONSTRAINT total_rooms >= occupied_rooms ENFORCED
func ChangeTotalRooms(db *sql.DB, newRoomCount int) (sql.Result, error) {
	query := "UPDATE hotel_man.hotel SET total_rooms=? WHERE id=1"
	res, err := db.Exec(query, newRoomCount)
	if err != nil {
		log.Panicln(err)
	}
	return res, nil
}

// CreateGuest TODO: requires CONSTRAINT occupied_rooms <= total_rooms ENFORCED, occupied_rooms >= 0 ENFORCED
func CreateGuest(db *sql.DB, guest *models.Guest) (sql.Result, error) {
	incrementQuery := "UPDATE hotel_man.hotel SET occupied_rooms = occupied_rooms + 1 WHERE id=1"
	selectQuery := "SELECT occupied_rooms from hotel_man.hotel WHERE id=1"
	insertQuery := "INSERT INTO hotel_man.guest (name, check_in_date, check_out_date, room_number, payment) VALUES (?, ?, ?, ?, ?)"

	tx, err := db.Begin()
	if err != nil {
		log.Panicln(err)
	}

	incrementRes, incrementErr := tx.Exec(incrementQuery)
	if incrementErr != nil {
		tx.Rollback()
		log.Panicln(err)
	}
	log.Println(incrementRes)

	selectRes := tx.QueryRow(selectQuery)

	var roomCount int
	log.Println(selectRes)
	scanErr := selectRes.Scan(&roomCount)
	if err != nil {
		tx.Rollback()
		log.Panicln(scanErr)
	}

	insertRes, insertErr := tx.Exec(insertQuery, guest.Name, guest.CheckInDate.String(), guest.CheckOutDate.String(), roomCount, guest.Payment)
	if err != nil {
		tx.Rollback()
		log.Panicln(insertErr)
	}
	log.Println(insertRes)

	commitErr := tx.Commit()
	if commitErr != nil {
		tx.Rollback()
		log.Panicln(commitErr)
	}
	return insertRes, nil
}

func GetGuestByID(db *sql.DB, id int) *sql.Row {
	q := "SELECT * from hotel_man.guest WHERE id=?"
	row := db.QueryRow(q, id)
	return row
}

func GetAllGuests(db *sql.DB) (*sql.Rows, error) {
	query := "SELECT * from hotel_man.guest"
	rows, err := db.Query(query)
	if err != nil {
		log.Panicln(err)
	}
	return rows, nil
}

func UpdateCheckOutDate(db *sql.DB, t time.Time, id int) (sql.Result, error) {
	q := "UPDATE hotel_man.guest SET guest.check_out_date=? WHERE id = ?"
	rows, err := db.Exec(q, t.String(), id)
	if err != nil {
		log.Panicln(err)
	}
	return rows, nil
}

func DeleteGuest(db *sql.DB, id int) (sql.Result, error) {
	decrementRoomNum := "UPDATE hotel_man.hotel SET occupied_rooms = occupied_rooms - 1 WHERE id=1"
	q := "DELETE FROM hotel_man.guest WHERE id = ?"

	tx, err := db.Begin()
	if err != nil {
		log.Panicln(err)
	}

	descrementRes, decrementErr := tx.Exec(decrementRoomNum)
	if decrementErr != nil {
		tx.Rollback()
		log.Panicln(err)
	}
	log.Println(descrementRes)

	res, err := db.Exec(q, id)
	if err != nil {
		tx.Rollback()
		log.Panicln(err)
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		tx.Rollback()
		log.Panicln(commitErr)
	}
	return res, nil
}
