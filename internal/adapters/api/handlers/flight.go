package handlers

import (
	"encoding/json"
	"net/http"
	"travel-backend/internal/core/domain/models"
	"travel-backend/internal/ports/api"
	"travel-backend/pkg/utils"

	"github.com/gorilla/mux"
)

// FlightHandler handles flight-related API requests
type FlightHandler struct {
	FlightService api.FlightService
}

// NewFlightHandler creates a new instance of FlightHandler
func NewFlightHandler(flightService api.FlightService) *FlightHandler {
	return &FlightHandler{FlightService: flightService}
}

// GetFlights handles GET /flights
func (h *FlightHandler) GetFlights(w http.ResponseWriter, r *http.Request) {
	flights, err := h.FlightService.GetAllFlights()
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, flights)
}

// GetFlightByID handles GET /flights/{id}
func (h *FlightHandler) GetFlightByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	flight, err := h.FlightService.GetFlightByID(id)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, flight)
}

// CreateFlight handles POST /flights
func (h *FlightHandler) CreateFlight(w http.ResponseWriter, r *http.Request) {
	var flight models.Flight
	if err := json.NewDecoder(r.Body).Decode(&flight); err != nil {
		utils.HandleError(w, err)
		return
	}
	err := h.FlightService.CreateFlight(&flight)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, flight)
}

// UpdateFlight handles PUT /flights/{id}
func (h *FlightHandler) UpdateFlight(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var flight models.Flight
	if err := json.NewDecoder(r.Body).Decode(&flight); err != nil {
		utils.HandleError(w, err)
		return
	}
	updatedFlight, err := h.FlightService.UpdateFlight(id, &flight)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, updatedFlight)
}

// DeleteFlight handles DELETE /flights/{id}
func (h *FlightHandler) DeleteFlight(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.FlightService.DeleteFlight(id); err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}
