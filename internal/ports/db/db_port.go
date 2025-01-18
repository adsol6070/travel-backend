package db

import (
	"time"
	"travel-backend/internal/core/domain/models"
)

type HotelRepository interface {
	GetAllHotels() ([]models.Hotel, error)
	GetHotelByID(id string) (*models.Hotel, error)
	CreateHotel(hotel *models.Hotel) error
	GetHotelBookings(hotelID string) ([]models.Booking, error) // Get all bookings for a specific hotel
}

type FlightRepository interface {
	GetAllFlights() ([]models.Flight, error)
	GetFlightByID(id string) (*models.Flight, error)
	CreateFlight(flight *models.Flight) error
	GetFlightBookings(flightID string) ([]models.Booking, error) // Get all bookings for a specific flight
	GetAvailableSeats(flightID string) ([]models.Seat, error)    // Get available seats for a flight
}

type BookingRepository interface {
	GetAllBookings() ([]models.Booking, error)
	GetBookingByID(id string) (*models.Booking, error)
	CreateBooking(booking *models.Booking) error
	UpdateBookingStatus(id string, status string) error
	GetBookingsByUserID(userID string) ([]models.Booking, error) // Get all bookings for a specific user
}

type PassengerRepository interface {
	GetAllPassengersByBookingID(bookingID string) ([]models.Passenger, error) // Get all passengers for a booking
	AddPassenger(passenger *models.Passenger) error                           // Add a passenger
	AssignSeat(passengerSeat *models.Seat) error                              // Assign a seat to a passenger
	GetPassengerSeats(bookingID string) ([]models.Seat, error)                // Get all seats assigned to passengers in a booking
	AssignMeal(passengerMeal *models.Meal) error                              // Assign a meal to a passenger
	GetPassengerMeals(bookingID string) ([]models.Meal, error)                // Get all meals assigned to passengers in a booking
}

type SeatRepository interface {
	GetSeatByID(seatID string) (*models.Seat, error)
	UpdateSeatAvailability(seatID string, isAvailable bool) error       // Update seat availability
	GetAvailableSeatsByFlightID(flightID string) ([]models.Seat, error) // Get available seats for a flight
	AssignSeatToPassenger(seatID, passengerID, bookingID string) error  // Assign a seat to a passenger
}

type MealRepository interface {
	GetAllMeals() ([]models.Meal, error)                       // Get all available meals
	GetMealsByClass(class string) ([]models.Meal, error)       // Get meals for a specific class
	AddMeal(meal *models.Meal) error                           // Add a new meal
	GetPassengerMeals(bookingID string) ([]models.Meal, error) // Get all meals assigned to passengers in a booking
	AssignMealToPassenger(passengerMeal *models.Meal) error    // Assign a meal to a passenger
}

type ReportRepository interface {
	GetBookingReport(startDate, endDate time.Time) ([]models.Booking, error) // Generate a report of bookings within a date range
}
