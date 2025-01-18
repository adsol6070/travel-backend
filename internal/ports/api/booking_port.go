package api

import "travel-backend/internal/core/domain/models"

type BookingService interface {
	GetAllBookings() ([]models.Booking, error)
	GetBookingByID(id string) (*models.Booking, error)
	CreateBooking(booking *models.Booking) error
	UpdateBookingStatus(id string, status string) error
	GetBookingsByUserID(userID string) ([]models.Booking, error)
}
