package models

type Seat struct {
	SeatID      string `json:"seatID"`
	FlightID    string `json:"flightID"`
	SeatNumber  string `json:"seatNumber"`
	Class       string `json:"class"`
	IsAvailable bool   `json:"isAvailable"`
}
