package models

import "time"

type Flight struct {
	FlightID      string    `json:"flightID"`
	Airline       string    `json:"airline"`
	Origin        string    `json:"origin"`
	Destination   string    `json:"destination"`
	DepartureTime time.Time `json:"departureTime"`
	ArrivalTime   time.Time `json:"arrivalTime"`
	AircraftType  string    `json:"aircraftType"`
}
