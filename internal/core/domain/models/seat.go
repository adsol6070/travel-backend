package models

type Seat struct {
	SeatID      string `json:"seatID" dynamodbav:"seatID"`
	FlightID    string `json:"flightID" dynamodbav:"flightID"`
	SeatNumber  string `json:"seatNumber" dynamodbav:"seatNumber"`
	Class       string `json:"class" dynamodbav:"class"`
	IsAvailable bool   `json:"isAvailable" dynamodbav:"isAvailable"`
}
