package handlers

import (
	"encoding/json"
	"net/http"
	"travel-backend/internal/core/domain/models"
	"travel-backend/internal/ports/api"
	"travel-backend/pkg/utils"

	"github.com/gorilla/mux"
)

// BookingHandler handles booking-related API requests
type BookingHandler struct {
	BookingService api.BookingService
}

// NewBookingHandler creates a new instance of BookingHandler
func NewBookingHandler(bookingService api.BookingService) *BookingHandler {
	return &BookingHandler{BookingService: bookingService}
}

// GetBookings handles GET /bookings
func (h *BookingHandler) GetBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.BookingService.GetAllBookings()
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, bookings)
}

// GetBookingByID handles GET /bookings/{id}
func (h *BookingHandler) GetBookingByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	booking, err := h.BookingService.GetBookingByID(id)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, booking)
}

// CreateBooking handles POST /bookings
func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		utils.HandleError(w, err)
		return
	}
	err := h.BookingService.CreateBooking(&booking)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, booking)
}

// UpdateBooking handles PUT /bookings/{id}
func (h *BookingHandler) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		utils.HandleError(w, err)
		return
	}
	updatedBooking, err := h.BookingService.UpdateBooking(id, &booking)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, updatedBooking)
}

// UpdateBookingStatus handles PUT /bookings/{id}/status
func (h *BookingHandler) UpdateBookingStatus(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var statusRequest struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&statusRequest); err != nil {
		utils.HandleError(w, err)
		return
	}
	err := h.BookingService.UpdateBookingStatus(id, statusRequest.Status)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, nil)
}

// DeleteBooking handles DELETE /bookings/{id}
func (h *BookingHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.BookingService.DeleteBooking(id)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}

// GetBookingsByUserID handles GET /bookings/user/{userID}
func (h *BookingHandler) GetBookingsByUserID(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]
	bookings, err := h.BookingService.GetBookingsByUserID(userID)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, bookings)
}


