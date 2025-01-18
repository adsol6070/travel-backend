package api

import "travel-backend/internal/core/domain/models"

type FlightService interface {
	GetAllFlights() ([]models.Flight, error)
	GetFlightByID(id string) (*models.Flight, error)
	CreateFlight(flight *models.Flight) error
	GetFlightBookings(flightID string) ([]models.Booking, error)
	GetAvailableSeats(flightID string) ([]models.Seat, error)
}
