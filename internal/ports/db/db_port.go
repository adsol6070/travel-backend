package db

import (
	"time"
	"travel-backend/internal/core/domain/models"
)

type HotelRepository interface {
	GetAllHotels() ([]models.Hotel, error)
	GetHotelByID(id string) (*models.Hotel, error)
	CreateHotel(hotel *models.Hotel) error
	GetHotelBookings(hotelID string) ([]models.Booking, error)
}

type FlightRepository interface {
	GetAllFlights() ([]models.Flight, error)
	GetFlightByID(id string) (*models.Flight, error)
	CreateFlight(flight *models.Flight) error
	GetFlightBookings(flightID string) ([]models.Booking, error)
	GetAvailableSeats(flightID string) ([]models.Seat, error)
}

type BookingRepository interface {
	GetAllBookings() ([]models.Booking, error)
	GetBookingByID(id string) (*models.Booking, error)
	CreateBooking(booking *models.Booking) error
	UpdateBookingStatus(id string, status string) error
	GetBookingsByUserID(userID string) ([]models.Booking, error)
}

type PassengerRepository interface {
	GetAllPassengersByBookingID(bookingID string) ([]models.Passenger, error)
	AddPassenger(passenger *models.Passenger) error
	AssignSeat(passengerSeat *models.Seat) error
	GetPassengerSeats(bookingID string) ([]models.Seat, error)
	AssignMeal(passengerMeal *models.Meal) error
	GetPassengerMeals(bookingID string) ([]models.Meal, error)
}

type SeatRepository interface {
	GetSeatByID(seatID string) (*models.Seat, error)
	UpdateSeatAvailability(seatID string, isAvailable bool) error
	GetAvailableSeatsByFlightID(flightID string) ([]models.Seat, error)
	AssignSeatToPassenger(seatID, passengerID, bookingID string) error
}

type MealRepository interface {
	GetAllMeals() ([]models.Meal, error)
	GetMealsByClass(class string) ([]models.Meal, error)
	AddMeal(meal *models.Meal) error
	GetPassengerMeals(bookingID string) ([]models.Meal, error)
	AssignMealToPassenger(passengerMeal *models.Meal) error
}

type ReportRepository interface {
	GetBookingReport(startDate, endDate time.Time) ([]models.Booking, error)
}
