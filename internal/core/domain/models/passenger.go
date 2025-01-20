package models

type Passenger struct {
	PassengerID    string `json:"passengerID" dynamodbav:"passengerID"`
	BookingID      string `json:"bookingID" dynamodbav:"bookingID"`
	Name           string `json:"name" dynamodbav:"name"`
	Age            int    `json:"age" dynamodbav:"age"`
	Gender         string `json:"gender" dynamodbav:"gender"`
	PassportNumber string `json:"passportNumber" dynamodbav:"passportNumber"`
}
