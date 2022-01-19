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
	decrementRoomNum := "UPDATE hotel_man.guest SET occupied_rooms = occupied_rooms - 1 WHERE id=?"
	q := "DELETE FROM hotel_man.guest WHERE id = ?"
	tx, err := db.Begin()
	if err != nil {
		log.Panicln(err)
	}
	descrementRes, decrementErr := tx.Exec(decrementRoomNum)
	if decrementErr != nil {
		tx.RollBack()
		log.Panicln(err)
	}
	log.Println(descrementRes)
	res, err := tx.Exec(q, id)
	if err != nil {
		log.Panicln(err)
	}
	commitErr := tx.Commit()
	if commitErr != nil {
		tx.Rollback()
		log.Panicln(commitErr)
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
