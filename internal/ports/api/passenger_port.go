package api

import "travel-backend/internal/core/domain/models"

type PassengerService interface {
	GetAllPassengersByBookingID(bookingID string) ([]models.Passenger, error)
	AddPassenger(passenger *models.Passenger) error
	AssignSeat(passengerSeat *models.Seat) error
	GetPassengerSeats(bookingID string) ([]models.Seat, error)
	AssignMeal(passengerMeal *models.Meal) error
	GetPassengerMeals(bookingID string) ([]models.Meal, error)
}
