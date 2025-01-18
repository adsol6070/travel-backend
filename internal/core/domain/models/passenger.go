package models

type Passenger struct {
	PassengerID    string `json:"passengerID"`
	BookingID      string `json:"bookingID"`
	Name           string `json:"name"`
	Age            int    `json:"age"`
	Gender         string `json:"gender"`
	PassportNumber string `json:"passportNumber"`
}
