package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"travel-backend/internal/core/domain/models"
	"travel-backend/internal/ports/api"
	"travel-backend/pkg/utils"
)

// HotelHandler handles hotel-related API requests
type HotelHandler struct {
	HotelService api.HotelService
}

// NewHotelHandler creates a new instance of HotelHandler
func NewHotelHandler(hotelService api.HotelService) *HotelHandler {
	return &HotelHandler{HotelService: hotelService}
}

// GetHotels handles GET /hotels
func (h *HotelHandler) GetHotels(w http.ResponseWriter, r *http.Request) {
	hotels, err := h.HotelService.GetAllHotels()
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, hotels)
}

// GetHotelByID handles GET /hotels/{id}
func (h *HotelHandler) GetHotelByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	hotel, err := h.HotelService.GetHotelByID(id)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, hotel)
}

// CreateHotel handles POST /hotels
func (h *HotelHandler) CreateHotel(w http.ResponseWriter, r *http.Request) {
	var hotel models.Hotel
	if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
		utils.HandleError(w, err)
		return
	}
	err := h.HotelService.CreateHotel(&hotel)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, hotel)
}

// UpdateHotel handles PUT /hotels/{id}
func (h *HotelHandler) UpdateHotel(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var hotel models.Hotel
	if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
		utils.HandleError(w, err)
		return
	}
	updatedHotel, err := h.HotelService.UpdateHotel(id, &hotel)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, updatedHotel)
}

// DeleteHotel handles DELETE /hotels/{id}
func (h *HotelHandler) DeleteHotel(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.HotelService.DeleteHotel(id); err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}
