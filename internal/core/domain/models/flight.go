package models

import "time"

type Flight struct {
	FlightID      string    `json:"flightID" dynamodbav:"flightID"`
	Airline       string    `json:"airline" dynamodbav:"airline"`
	Origin        string    `json:"origin" dynamodbav:"origin"`
	Destination   string    `json:"destination" dynamodbav:"destination"`
	DepartureTime time.Time `json:"departureTime" dynamodbav:"departureTime"`
	ArrivalTime   time.Time `json:"arrivalTime" dynamodbav:"arrivalTime"`
	AircraftType  string    `json:"aircraftType" dynamodbav:"aircraftType"`
}
