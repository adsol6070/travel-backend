package api

import "travel-backend/internal/core/domain/models"

type HotelService interface {
	GetAllHotels() ([]models.Hotel, error)
	GetHotelByID(id string) (*models.Hotel, error)
	CreateHotel(hotel *models.Hotel) error
	GetHotelBookings(hotelID string) ([]models.Booking, error)
	UpdateHotel(id string, hotel *models.Hotel) (*models.Hotel, error)
	DeleteHotel(id string) error
}
