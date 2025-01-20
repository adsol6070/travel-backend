package api

import (
	"net/http"
	"travel-backend/internal/adapters/api/handlers"

	"github.com/gorilla/mux"
)

// SetupRoutes sets up the API routes
func SetupRoutes(router *mux.Router, hotelHandler *handlers.HotelHandler, flightHandler *handlers.FlightHandler, bookingHandler *handlers.BookingHandler) {
	// Hotel routes
	hotelRouter := router.PathPrefix("/hotels").Subrouter()
	hotelRouter.HandleFunc("/", hotelHandler.GetHotels).Methods(http.MethodGet)
	hotelRouter.HandleFunc("/{id}", hotelHandler.GetHotelByID).Methods(http.MethodGet)
	hotelRouter.HandleFunc("/", hotelHandler.CreateHotel).Methods(http.MethodPost)
	hotelRouter.HandleFunc("/{id}", hotelHandler.UpdateHotel).Methods(http.MethodPut)
	hotelRouter.HandleFunc("/{id}", hotelHandler.DeleteHotel).Methods(http.MethodDelete)

	// Flight routes
	flightRouter := router.PathPrefix("/flights").Subrouter()
	flightRouter.HandleFunc("/", flightHandler.GetFlights).Methods(http.MethodGet)
	flightRouter.HandleFunc("/{id}", flightHandler.GetFlightByID).Methods(http.MethodGet)
	flightRouter.HandleFunc("/", flightHandler.CreateFlight).Methods(http.MethodPost)
	flightRouter.HandleFunc("/{id}", flightHandler.UpdateFlight).Methods(http.MethodPut)
	flightRouter.HandleFunc("/{id}", flightHandler.DeleteFlight).Methods(http.MethodDelete)

	// Booking routes
	bookingRouter := router.PathPrefix("/bookings").Subrouter()
	bookingRouter.HandleFunc("/", bookingHandler.CreateBooking).Methods(http.MethodPost)
	bookingRouter.HandleFunc("/{id}", bookingHandler.GetBookingByID).Methods(http.MethodGet)
	bookingRouter.HandleFunc("/{id}", bookingHandler.UpdateBooking).Methods(http.MethodPut)
	bookingRouter.HandleFunc("/{id}", bookingHandler.DeleteBooking).Methods(http.MethodDelete)
}
