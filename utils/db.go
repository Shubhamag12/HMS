package utils

import (
	"database/sql"
	"log"

	"github.com/Shubhamag12/HMS/models"
)

func CreateHotel(db *sql.DB, hotel *models.Hotel) (sql.Result, error) {
	query := "INSERT INTO hotel_man.hotel (name, room_count, unoccupied_rooms, cost_per_day) VALUES (?, ?, ?, ?, ?)"
	res, err := db.Exec(query, hotel.Name, hotel.RoomCount, hotel.UnoccupiedRooms, hotel.CostPerDay)
	if err != nil {
		// TODO: better error handling
		log.Fatal(err)
	}
	log.Println(res)
	return res, nil
}

func CreateGuest(db *sql.DB, guest *models.Guest, payment *models.Payment) (sql.Result, sql.Result, error) {
	// TODO: execute query to do the payment here
	paymentRes, err := CreatePayment(db, payment)
	if err != nil {
		log.Panic(err)
		panic(err)
	}
	log.Println(paymentRes)
	query := "INSERT INTO hotel_man.guest (hotel_id, payment_id, name, check_in_date, check_out_date, room_number) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := db.Exec(query, guest.HotelId, guest.PaymentId, guest.Name, guest.CheckInDate, guest.CheckOutDate, guest.RoomNumber)
	if err != nil {
		log.Panic(err)
		panic(err)
	}
	return res, paymentRes, nil
}

func CreatePayment(db *sql.DB, payment *models.Payment) (sql.Result, error) {
	query := "INSERT INTO hotel_man.payment (booking_payment, is_checkout_extended, overdue) VALUES (?, ?, ?)"
	res, err := db.Exec(query, payment.BookingPayment, payment.IsCheckoutExtended, payment.Overdue)
	if err != nil {
		log.Panic(err)
		panic(err)
	}
	return res, nil
}
