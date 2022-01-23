package utils

import (
	"database/sql"
	"errors"
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
		return nil, err
	}
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
		return nil, err
	}
	return res, nil
}

// CreateGuest TODO: requires CONSTRAINT occupied_rooms <= total_rooms ENFORCED, occupied_rooms >= 0 ENFORCED
func CreateGuest(db *sql.DB, guest *models.Guest) (sql.Result, error) {
	incrementQuery := "UPDATE hotel_man.hotel SET occupied_rooms = occupied_rooms + 1 WHERE id=1"
	selectQuery := "SELECT cost_per_day from hotel_man.hotel WHERE id=1"
	insertQuery := "INSERT INTO hotel_man.guest (name, check_in_date, check_out_date, payment) VALUES (?, ?, ?, ?)"

	checkInDate, _ := time.Parse("2006-01-02", guest.CheckInDate)
	checkOutDate, _ := time.Parse("2006-01-02", guest.CheckOutDate)
	delta := int(checkOutDate.Sub(checkInDate).Hours() / 24)

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	incrementRes, incrementErr := tx.Exec(incrementQuery)
	if incrementErr != nil {
		tx.Rollback()
		return nil, incrementErr
	}
	log.Println(incrementRes.LastInsertId())

	selectRes := tx.QueryRow(selectQuery)

	var costPerDay int
	log.Println(selectRes)
	scanErr := selectRes.Scan(&costPerDay)
	if err != nil {
		tx.Rollback()
		return nil, scanErr
	}
	guest.Payment = costPerDay * delta

	insertRes, insertErr := tx.Exec(insertQuery, guest.Name, guest.CheckInDate, guest.CheckOutDate, guest.Payment)
	if err != nil {
		tx.Rollback()
		return nil, insertErr
	}
	log.Println(insertRes)

	commitErr := tx.Commit()
	if commitErr != nil {
		tx.Rollback()
		return nil, commitErr
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
		return nil, err
	}
	return rows, nil
}

func UpdateCheckOutDate(db *sql.DB, t time.Time, id int) (sql.Result, error) {
	checkInQuery := "SELECT check_in_date from hotel_man.guest WHERE id = ?"
	costPerDayQuery := "SELECT cost_per_day from hotel_man.hotel WHERE id=1"
	updateQuery := "UPDATE hotel_man.guest SET check_out_date=?, payment=? WHERE id = ?"
	var costPerDay int
	var checkInDateString string
	checkInRes := db.QueryRow(checkInQuery, id)
	costPerDayRes := db.QueryRow(costPerDayQuery)

	checkInRes.Scan(&checkInDateString)
	costPerDayRes.Scan(&costPerDay)

	checkInDate, _ := time.Parse("2006-01-02", checkInDateString)
	delta := int(t.Sub(checkInDate).Hours() / 24)
	if delta <= 0 {
		return nil, errors.New("invalid date range")
	}
	newPaymentAmount := costPerDay * delta
	rows, err := db.Exec(updateQuery, t.Format("2006-01-02"), newPaymentAmount, id)
	if err != nil {
		return nil,err
	}
	return rows, nil
}

func DeleteGuest(db *sql.DB, id int) (sql.Result, error) {
	decrementRoomNum := "UPDATE hotel_man.hotel SET occupied_rooms = occupied_rooms - 1 WHERE id=1"
	q := "DELETE FROM hotel_man.guest WHERE id = ?"

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	descrementRes, decrementErr := tx.Exec(decrementRoomNum)
	if decrementErr != nil {
		tx.Rollback()
		return nil, decrementErr
	}
	log.Println(descrementRes)

	res, err := db.Exec(q, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		tx.Rollback()
		return nil, commitErr
	}
	return res, nil
}
