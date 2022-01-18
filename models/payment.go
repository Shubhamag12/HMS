package models

type Payment struct {
	Id                 int
	BookingPayment     int
	IsCheckoutExtended bool
	Overdue            int
}

// id: PK
// bookingPayment
// isCheckoutExtended
// overdue
