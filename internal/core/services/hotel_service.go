package services

import (
	"errors"
	"travel-backend/internal/core/domain/models"
	"travel-backend/internal/ports/db"
)

type HotelServiceImpl struct {
	hotelRepo   db.HotelRepository
	bookingRepo db.BookingRepository
}

// NewHotelService creates a new instance of HotelServiceImpl
func NewHotelService(hotelRepo db.HotelRepository, bookingRepo db.BookingRepository) *HotelServiceImpl {
	return &HotelServiceImpl{
		hotelRepo:   hotelRepo,
		bookingRepo: bookingRepo,
	}
}

// GetAllHotels retrieves all hotels from the repository
func (s *HotelServiceImpl) GetAllHotels() ([]models.Hotel, error) {
	return s.hotelRepo.GetAllHotels()
}

// GetHotelByID retrieves a hotel by its ID
func (s *HotelServiceImpl) GetHotelByID(id string) (*models.Hotel, error) {
	if id == "" {
		return nil, errors.New("hotel ID cannot be empty")
	}
	return s.hotelRepo.GetHotelByID(id)
}

// CreateHotel creates a new hotel in the repository
func (s *HotelServiceImpl) CreateHotel(hotel *models.Hotel) error {
	if hotel == nil {
		return errors.New("hotel details cannot be nil")
	}
	if hotel.HotelID == "" {
		return errors.New("hotel ID is required")
	}
	return s.hotelRepo.CreateHotel(hotel)
}

// GetHotelBookings retrieves all bookings for a specific hotel
func (s *HotelServiceImpl) GetHotelBookings(hotelID string) ([]models.Booking, error) {
	if hotelID == "" {
		return nil, errors.New("hotel ID cannot be empty")
	}

	// Assuming bookingRepo has a method to get bookings by hotelID
	return s.bookingRepo.GetBookingsByUserID(hotelID)
}

// UpdateHotel updates a hotel's details
func (s *HotelServiceImpl) UpdateHotel(id string, hotel *models.Hotel) (*models.Hotel, error) {
	if id == "" {
		return nil, errors.New("hotel ID cannot be empty")
	}
	if hotel == nil {
		return nil, errors.New("hotel details cannot be nil")
	}

	// Assuming hotelRepo has a method to update a hotel
	updatedHotel, err := s.hotelRepo.UpdateHotel(id, hotel)
	if err != nil {
		return nil, err
	}

	return updatedHotel, nil
}

// DeleteHotel deletes a hotel by its ID
func (s *HotelServiceImpl) DeleteHotel(id string) error {
	if id == "" {
		return errors.New("hotel ID cannot be empty")
	}

	// Assuming hotelRepo has a method to delete a hotel
	err := s.hotelRepo.DeleteHotel(id)
	if err != nil {
		return err
	}

	return nil
}
