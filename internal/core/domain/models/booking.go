package models

import "time"

type Booking struct {
	BookingID     string    `json:"bookingID" dynamodbav:"bookingID"`
	UserID        string    `json:"userID" dynamodbav:"userID"`
	FlightID      string    `json:"flightID" dynamodbav:"flightID"`
	BookingStatus string    `json:"bookingStatus" dynamodbav:"bookingStatus"`
	CreatedAt     time.Time `json:"createdAt" dynamodbav:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt" dynamodbav:"updatedAt"`
}
