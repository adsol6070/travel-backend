package services

import (
	"errors"
	"travel-backend/internal/core/domain/models"
	"travel-backend/internal/ports/db"
)

type BookingServiceImpl struct {
	bookingRepo db.BookingRepository
}

func NewBookingService(bookingRepo db.BookingRepository) *BookingServiceImpl {
	return &BookingServiceImpl{
		bookingRepo: bookingRepo,
	}
}

func (s *BookingServiceImpl) GetAllBookings() ([]models.Booking, error) {
	bookings, err := s.bookingRepo.GetAllBookings()
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (s *BookingServiceImpl) GetBookingByID(id string) (*models.Booking, error) {
	booking, err := s.bookingRepo.GetBookingByID(id)
	if err != nil {
		return nil, err
	}
	if booking == nil {
		return nil, errors.New("booking not found")
	}
	return booking, nil
}

func (s *BookingServiceImpl) CreateBooking(booking *models.Booking) error {
	if booking == nil {
		return errors.New("invalid booking details")
	}
	err := s.bookingRepo.CreateBooking(booking)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookingServiceImpl) UpdateBookingStatus(id string, status string) error {
	if id == "" || status == "" {
		return errors.New("invalid booking ID or status")
	}
	err := s.bookingRepo.UpdateBookingStatus(id, status)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookingServiceImpl) GetBookingsByUserID(userID string) ([]models.Booking, error) {
	if userID == "" {
		return nil, errors.New("invalid user ID")
	}
	bookings, err := s.bookingRepo.GetBookingsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (s *BookingServiceImpl) DeleteBooking(id string) error {
	if id == "" {
		return errors.New("invalid booking ID")
	}

	// Ensure the booking exists before attempting to delete it
	booking, err := s.bookingRepo.GetBookingByID(id)
	if err != nil {
		return err
	}
	if booking == nil {
		return errors.New("booking not found")
	}

	// Call the repository to delete the booking
	err = s.bookingRepo.DeleteBooking(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *BookingServiceImpl) UpdateBooking(id string, booking *models.Booking) (*models.Booking, error) {
	if id == "" {
		return nil, errors.New("invalid booking ID")
	}
	if booking == nil {
		return nil, errors.New("invalid booking details")
	}

	// Ensure the booking exists before attempting to update it
	existingBooking, err := s.bookingRepo.GetBookingByID(id)
	if err != nil {
		return nil, err
	}
	if existingBooking == nil {
		return nil, errors.New("booking not found")
	}

	// Call the repository to update the booking
	updatedBooking, err := s.bookingRepo.UpdateBooking(id, booking)
	if err != nil {
		return nil, err
	}

	return updatedBooking, nil
}
