package main

import (
	"fmt"
	"log"
	"net/http"
	"travel-backend/customConfig"
	"travel-backend/internal/adapters/api"
	"travel-backend/internal/adapters/api/handlers"
	"travel-backend/internal/adapters/db/dynamodb"
	"travel-backend/internal/core/services"

	"github.com/gorilla/mux"
)

func main() {
	customConfig.LoadConfig()

	dbClient := dynamodb.NewDynamoDBClient()

	// Initialize repositories
	hotelRepo := dynamodb.NewHotelRepo(dbClient)
	flightRepo := dynamodb.NewFlightRepo(dbClient)
	bookingRepo := dynamodb.NewBookingRepo(dbClient)

	// Initialize services
	hotelService := services.NewHotelService(hotelRepo, bookingRepo)
	flightService := services.NewFlightService(flightRepo, bookingRepo)
	bookingService := services.NewBookingService(bookingRepo)

	// Initialize API Handlers
	hotelHandler := handlers.NewHotelHandler(hotelService)
	flightHandler := handlers.NewFlightHandler(flightService)
	bookingHandler := handlers.NewBookingHandler(bookingService)

	// Set up routes
	router := mux.NewRouter()
	api.SetupRoutes(router, hotelHandler, flightHandler, bookingHandler)

	// Start the server
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("Server started on http://localhost:3000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}

}
