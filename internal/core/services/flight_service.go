package services

import (
	"errors"
	"travel-backend/internal/core/domain/models"
	"travel-backend/internal/ports/db"
)

type FlightServiceImpl struct {
	flightRepo  db.FlightRepository
	bookingRepo db.BookingRepository
	seatRepo    db.SeatRepository
}

func NewFlightService(flightRepo db.FlightRepository, bookingRepo db.BookingRepository, /* seatRepo db.SeatRepository */) *FlightServiceImpl {
	return &FlightServiceImpl{
		flightRepo:  flightRepo,
		bookingRepo: bookingRepo,
		// seatRepo:    seatRepo,
	}
}

func (s *FlightServiceImpl) GetAllFlights() ([]models.Flight, error) {
	return s.flightRepo.GetAllFlights()
}

func (s *FlightServiceImpl) GetFlightByID(id string) (*models.Flight, error) {
	if id == "" {
		return nil, errors.New("flight ID cannot be empty")
	}
	return s.flightRepo.GetFlightByID(id)
}

func (s *FlightServiceImpl) CreateFlight(flight *models.Flight) error {
	if flight == nil {
		return errors.New("flight details cannot be nil")
	}
	if flight.FlightID == "" || flight.Airline == "" {
		return errors.New("flight ID and airline are required")
	}
	return s.flightRepo.CreateFlight(flight)
}

// GetFlightBookings retrieves all bookings for a specific flight
func (s *FlightServiceImpl) GetFlightBookings(flightID string) ([]models.Booking, error) {
	if flightID == "" {
		return nil, errors.New("flight ID cannot be empty")
	}
	return s.flightRepo.GetFlightBookings(flightID)
}

// GetAvailableSeats retrieves all available seats for a specific flight
func (s *FlightServiceImpl) GetAvailableSeats(flightID string) ([]models.Seat, error) {
	// if flightID == "" {
	// 	return nil, errors.New("flight ID cannot be empty")
	// }
	// seats, err := s.seatRepo.GetAvailableSeatsByFlightID(flightID)
	// if err != nil {
	// 	return nil, err
	// }

	// Filter available seats
	var availableSeats []models.Seat
	// for _, seat := range seats {
	// 	if seat.IsAvailable {
	// 		availableSeats = append(availableSeats, seat)
	// 	}
	// }

	return availableSeats, nil
}

// UpdateFlight updates the details of a flight by its ID
func (s *FlightServiceImpl) UpdateFlight(id string, flight *models.Flight) (*models.Flight, error) {
	if id == "" {
		return nil, errors.New("flight ID cannot be empty")
	}
	if flight == nil {
		return nil, errors.New("flight details cannot be nil")
	}

	// Fetch the existing flight to ensure it exists
	existingFlight, err := s.flightRepo.GetFlightByID(id)
	if err != nil {
		return nil, err
	}
	if existingFlight == nil {
		return nil, errors.New("flight not found")
	}

	// Update the flight details
	updatedFlight, err := s.flightRepo.UpdateFlight(id, flight)
	if err != nil {
		return nil, err
	}

	return updatedFlight, nil
}

// DeleteFlight deletes a flight by its ID
func (s *FlightServiceImpl) DeleteFlight(id string) error {
	if id == "" {
		return errors.New("flight ID cannot be empty")
	}

	// Fetch the existing flight to ensure it exists
	existingFlight, err := s.flightRepo.GetFlightByID(id)
	if err != nil {
		return err
	}
	if existingFlight == nil {
		return errors.New("flight not found")
	}

	// Check if there are any bookings associated with the flight
	bookings, err := s.flightRepo.GetFlightBookings(id)
	if err != nil {
		return err
	}
	if len(bookings) > 0 {
		return errors.New("flight has active bookings and cannot be deleted")
	}

	// Delete the flight
	errr := s.flightRepo.DeleteFlight(id)
	if errr != nil {
		return errr
	}

	return nil
}
