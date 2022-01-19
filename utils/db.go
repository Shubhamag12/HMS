package utils

import (
	"database/sql"
	"log"
	"time"
)

func GetGuestsByID(db *sql.DB, id int) *sql.Row {
	q := "SELECT * from hotel_man.guest WHERE id=?"
	rows := db.QueryRow(q, id)
	return rows
}

func DeleteGuest(db *sql.DB, id int) (sql.Result, error) {
	q := "DELETE FROM hotel_man.guest WHERE id = ?"
	res, err := db.Exec(q, id)
	if err != nil {
		log.Panicln(err)
	}
	return res, nil
}

func UpdateCheckOutDate(db *sql.DB, t time.Time, id int) (sql.Result, error) {
	q := "UPDATE hotel_man.guest SET guest.check_out_date=? WHERE id = ?"
	rows, err := db.Exec(q, t.String(), id)
	if err != nil {
		log.Panicln(err)
	}
	return rows, nil
}
