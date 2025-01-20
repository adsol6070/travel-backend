package models

import "time"

type Asset struct {
	Type   string `json:"type"`
	S3Link string `json:"s3Link"`
}

type Hotel struct {
	HotelID           string    `json:"hotelID" dynamodbav:"hotelID"`
	BookingID         string    `json:"bookingID" dynamodbav:"bookingID"`
	UserID            string    `json:"userID" dynamodbav:"userID"`
	CheckInDate       time.Time `json:"checkInDate" dynamodbav:"checkInDate"`
	CheckOutDate      time.Time `json:"checkOutDate" dynamodbav:"checkOutDate"`
	IsCheckinFlexible bool      `json:"isCheckinFlexible" dynamodbav:"isCheckinFlexible"`
	BookingStatus     string    `json:"bookingStatus" dynamodbav:"bookingStatus"`
	CreatedAt         time.Time `json:"createdAt" dynamodbav:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt" dynamodbav:"updatedAt"`
	PaymentStatus     string    `json:"paymentStatus" dynamodbav:"paymentStatus"`
	RoomType          string    `json:"roomType" dynamodbav:"roomType"`
	NumberOfGuests    int       `json:"numberOfGuests" dynamodbav:"numberOfGuests"`
	SpecialRequests   string    `json:"specialRequests" dynamodbav:"specialRequests"`
	Assets            []Asset   `json:"assets" dynamodbav:"assets"`
}
